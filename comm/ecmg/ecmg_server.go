package ecmg

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ranger2011/cas-go/comm"
	"github.com/ranger2011/cas-go/comm/data"
	"github.com/ranger2011/cas-go/utils"
)

const maxBufferSize int = 1024
const ecmTableId byte = 0x51

func StartEcmg(port uint16) bool {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return false
	}
	defer listener.Close()

	go handleAccept(listener)
	return true
}

func handleAccept(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buffer := make([]byte, maxBufferSize)
	reader := bufio.NewReader(conn)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			conn.Close()
			return
		}
		var message comm.Message
		if !message.Parse(buffer, n) {
			conn.Close()
			return
		}

		switch message.Header.Type {
		case ECM_MESSAGE_TYPE_CHANNEL_SETUP:
			processChannelSetup(&message, conn)
		case ECM_MESSAGE_TYPE_CHANNEL_TEST:
			processChannelTest(&message, conn)
		case ECM_MESSAGE_TYPE_CHANNEL_CLOSE:
			conn.Close()
			return
		case ECM_MESSAGE_TYPE_CHANNEL_ERROR:
			conn.Close()
			return //we don't process error, just close the connection
		case ECM_MESSAGE_TYPE_STREAM_SETUP:
			processStreamSetupOrTest(&message, conn)
		case ECM_MESSAGE_TYPE_STREAM_TEST:
			processStreamSetupOrTest(&message, conn)
		case ECM_MESSAGE_TYPE_STREAM_CLOSE_REQUEST:
			processStreamCloseRequest(&message, conn)
		case ECM_MESSAGE_TYPE_STREAM_ERROR:
			conn.Close()
			return
		case ECM_MESSAGE_TYPE_CW_PROVISION:
			processCwProvision(&message, conn)
		default:
			fmt.Printf("Unknown message type: %d.\n", message.Header.Type)
		}
	}
}

func processChannelSetup(message *comm.Message, conn net.Conn) {
	found, superCasId := GetSuperCasId(message)
	if !found {
		return
	}
	if superCasId != comm.SuperCasId {
		fmt.Println("Error super cas id!")
		return
	}

	found, channelId := GetDataChannelId(message)
	if !found {
		return
	}
	buffer := make([]byte, maxBufferSize)
	pkgLen := ChannelSetup(channelId, superCasId, buffer)
	utils.SendTcp(buffer, pkgLen, conn)
}

func processChannelTest(message *comm.Message, conn net.Conn) {
	found, channelId := GetDataChannelId(message)
	if !found {
		return
	}

	buffer := make([]byte, maxBufferSize)
	pkgLen := ChannelStatus(channelId, buffer)
	utils.SendTcp(buffer, pkgLen, conn)
}

func processStreamSetupOrTest(message *comm.Message, conn net.Conn) {
	found, channelId := GetDataChannelId(message)
	if !found {
		return
	}

	found, streamId := GetDataStreamId(message)
	if !found {
		return
	}

	found, ecmId := GetEcmId(message)
	if !found {
		return
	}

	buffer := make([]byte, maxBufferSize)
	pkgLen := StreamStatus(channelId, streamId, ecmId, ACCESS_CRITERIA_TRANSFER_MODE, buffer)
	utils.SendTcp(buffer, pkgLen, conn)
}

func processStreamCloseRequest(message *comm.Message, conn net.Conn) {
	found, channelId := GetDataChannelId(message)
	if !found {
		return
	}

	found, streamId := GetDataStreamId(message)
	if !found {
		return
	}

	buffer := make([]byte, maxBufferSize)
	pkgLen := StreamCloseResponse(channelId, streamId, buffer)
	utils.SendTcp(buffer, pkgLen, conn)
}

//here we send some plain data, just for demo
//for product code, you must include channel info, service key, and time etc.., encrypt data, add timestamp etc..
func generateEcmTable(
	version byte,
	superCasId uint32,
	channelId uint16,
	channelIndex uint16,
	ppvCode uint16,
	packageNumber byte,
	packageIdList []uint16,
	year uint16,
	month uint16,
	day uint16,
	hour uint16,
	minute uint16,
	second uint16,
	parity bool,
	serviceKey []byte,
	cw []byte,
	buffer []byte,
) uint16 {
	buffer[0] = ecmTableId
	buffer[1] = version
	byteCount := 3
	byteCount += utils.PackageUint32(buffer, byteCount, superCasId)
	byteCount += utils.PackageUint16(buffer, byteCount, channelId)
	byteCount += utils.PackageUint16(buffer, byteCount, channelIndex)
	byteCount += utils.PackageUint16(buffer, byteCount, ppvCode)
	byteCount += utils.PackageUint16(buffer, byteCount, year)
	buffer[byteCount] = byte(month)
	byteCount++
	buffer[byteCount] = byte(day)
	byteCount++
	buffer[byteCount] = byte(hour)
	byteCount++
	buffer[byteCount] = byte(minute)
	byteCount++
	buffer[byteCount] = byte(second)
	if parity {
		buffer[byteCount] = 0
	} else {
		buffer[byteCount] = 1
	}
	byteCount++
	encrypted := utils.EncryptAes128Ecb(cw, serviceKey)
	utils.MemCpy(buffer, encrypted, byteCount, 0, 16)
	byteCount += 16
	buffer[byteCount] = packageNumber
	if packageNumber > 0 {
		var i byte
		for i = 0; i < packageNumber; i++ {
			byteCount += utils.PackageUint16(buffer, byteCount, packageIdList[i])
		}
	}
	buffer[byteCount] = packageNumber
	dataForCrc := make([]byte, byteCount-3)
	utils.MemCpy(dataForCrc, buffer, 0, 3, byteCount-3)
	crc := utils.CheckSumCCITT(dataForCrc)
	byteCount += utils.PackageUint16(buffer, byteCount, crc)
	buffer[2] = byte(byteCount - 3)
	return uint16(byteCount)
}

func processCwProvision(message *comm.Message, conn net.Conn) {
	found, channelId := GetDataChannelId(message)
	if !found {
		return
	}

	found, streamId := GetDataStreamId(message)
	if !found {
		return
	}

	found, cpNumber := GetCpNumber(message)
	if !found {
		return
	}

	found, cw := GetCw(message)
	if !found {
		return
	}

	found, ac := GetAccessCriteria(message)
	if !found {
		return
	}
	//we don't use ppv for now
	var ppvChannelIndex uint16 = 0xffff //uint16(ac[0]<<8) + uint16(ac[1])
	var ppvCode uint16 = 0
	var productId = uint16(ac[2])<<8 + uint16(ac[3])

	productIdArray := data.FindPackagesIncludeChannel(productId)

	currentTime := time.Now()

	ecm := make([]byte, 128)
	len := generateEcmTable(
		0,
		comm.SuperCasId,
		productId,
		ppvChannelIndex,
		ppvCode,
		byte(len(productIdArray)),
		productIdArray,
		uint16(currentTime.Year()),
		uint16(currentTime.Month()),
		uint16(currentTime.Day()),
		uint16(currentTime.Hour()),
		uint16(currentTime.Minute()),
		uint16(currentTime.Second()),
		data.IsCurrentKeyOdd(),
		data.GetCurrentKey(),
		cw,
		ecm,
	)
	buffer := make([]byte, maxBufferSize)
	pkgLen := EcmResponse(channelId, streamId, cpNumber, len, ecm, buffer)
	utils.SendTcp(buffer, pkgLen, conn)
}
