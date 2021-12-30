package ecmg

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/ranger2011/cas-go/mux"
)

const maxBufferSize int = 1024
const ecmTableId byte = 0x55

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
		var message mux.Message
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

func processChannelSetup(message *mux.Message, conn net.Conn) {
	found, superCasId := GetSuperCasId(message)
	if !found {
		return
	}
	if superCasId != mux.SuperCasId {
		fmt.Println("Error super cas id!")
		return
	}

	found, channelId := GetDataChannelId(message)
	if !found {
		return
	}
	buffer := make([]byte, maxBufferSize)
	pkgLen := ChannelSetup(channelId, superCasId, buffer)
	mux.SendTcp(buffer, pkgLen, conn)
}

func processChannelTest(message *mux.Message, conn net.Conn) {
	found, channelId := GetDataChannelId(message)
	if !found {
		return
	}

	buffer := make([]byte, maxBufferSize)
	pkgLen := ChannelStatus(channelId, buffer)
	mux.SendTcp(buffer, pkgLen, conn)
}

func processStreamSetupOrTest(message *mux.Message, conn net.Conn) {
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
	mux.SendTcp(buffer, pkgLen, conn)
}

func processStreamCloseRequest(message *mux.Message, conn net.Conn) {
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
	mux.SendTcp(buffer, pkgLen, conn)
}

//here we send some plain data, just for demo
//for product code, you must include channel info, service key, and time etc.., encrypt data, add timestamp etc..
func generateEcmTable(cw []byte) []byte {
	buffer := make([]byte, 7+len(cw))
	buffer[0] = ecmTableId
	buffer[1] = 0
	buffer[2] = byte(len(cw)) + 4
	mux.MemCpy(buffer, cw, 3, 0, len(cw))
	return buffer
}

func processCwProvision(message *mux.Message, conn net.Conn) {
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

	section := generateEcmTable(cw)

	buffer := make([]byte, maxBufferSize)
	pkgLen := EcmResponse(channelId, streamId, cpNumber, section, buffer)
	mux.SendTcp(buffer, pkgLen, conn)
}
