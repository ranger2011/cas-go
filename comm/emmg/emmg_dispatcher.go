package emmg

import (
	"fmt"
	"net"
	"time"

	"github.com/ranger2011/cas-go/comm"
	"github.com/ranger2011/cas-go/config"
	"github.com/ranger2011/cas-go/utils"
)

const channelId uint16 = 1
const streamId uint16 = 1
const dataId uint16 = 1
const dataType byte = 0
const maxBufferSize = 1024

func Dispatch(addr string, msg chan []byte) {
	for {
		conn, connected := installChannel(addr)
		if !connected {
			fmt.Println("Can't install channel and stream...")
			fmt.Println("5 seconds later try again...")
			time.Sleep(5 * time.Second)
			continue
		}
		for {
			data := <-msg
			if !utils.SendTcp(data, len(data), conn) {
				conn.Close()
				break
			}
		}
	}
}

func installChannel(addr string) (net.Conn, bool) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	buffer := make([]byte, maxBufferSize)
	pkgLen := ChannelSetup(comm.SuperCasId, channelId, config.GetDataMode(), buffer)
	utils.SendTcp(buffer, pkgLen, conn)
	response, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Can't receive channel install response for emmg...")
		conn.Close()
		return nil, false
	}
	var message comm.Message
	if !message.Parse(buffer, response) {
		conn.Close()
		fmt.Println("Can't parse emmg channel install response message.")
		return nil, false
	}
	if message.Header.Type != MessageEmmChannelStatus {
		fmt.Println("Error response message type %u for emmg channel setup.", message.Header.Type)
		_ = conn.Close()
		return nil, false
	}
	//for simplify, here we ignore super cas id validate...
	pkgLen = StreamSetup(comm.SuperCasId, channelId, streamId, dataId, dataType, buffer)
	utils.SendTcp(buffer, pkgLen, conn)

	response, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Can't receive stream install response for emmg...")
		conn.Close()
		return nil, false
	}
	if !message.Parse(buffer, response) {
		conn.Close()
		fmt.Println("Can't parse emmg stream install response message.")
		return nil, false
	}
	if message.Header.Type != MessageEmmStreamStatus {
		fmt.Println("Error response message type %u for emmg stream setup.", message.Header.Type)
		conn.Close()
		return nil, false
	}
	return conn, true
}
