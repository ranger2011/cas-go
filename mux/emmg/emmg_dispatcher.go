package emmg

import (
	"fmt"
	"net"
	"time"

	"github.com/ranger2011/cas-go/mux"
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
			if !mux.SendTcp(data, len(data), conn) {
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
	pkgLen := ChannelSetup(mux.SuperCasId, channelId, mux.SectionTspktFlagSection, buffer)
	mux.SendTcp(buffer, pkgLen, conn)
	response, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Can't receive channel install response for emmg...")
		conn.Close()
		return nil, false
	}
	var message mux.Message
	if !message.Parse(buffer, response) {
		conn.Close()
		fmt.Println("Can't parse emmg channel install response message.")
		return nil, false
	}
	if message.Header.Type != MESSAGE_EMM_CHANNEL_STATUS {
		fmt.Println("Error response message type %u for emmg channel setup.", message.Header.Type)
		conn.Close()
		return nil, false
	}
	//for simplify, here we ignore super cas id validate...
	pkgLen = StreamSetup(mux.SuperCasId, channelId, streamId, dataId, dataType, buffer)
	mux.SendTcp(buffer, pkgLen, conn)

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
	if message.Header.Type != MESSAGE_EMM_STREAM_STATUS {
		fmt.Println("Error response message type %u for emmg stream setup.", message.Header.Type)
		conn.Close()
		return nil, false
	}
	return conn, true
}
