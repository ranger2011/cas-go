package main

import (
	"strconv"

	"github.com/ranger2011/cas-go/comm/data"
	"github.com/ranger2011/cas-go/utils"
)

func main() {
	// fmt.Println("Conditional access system demo....")
	// fmt.Printf("The protocol verison is %d.\n", comm.ProtocolVersion)
	// //start a ecmg server
	// go ecmg.StartEcmg(5555)
	// emmChan := make(chan []byte)
	// //start a emmg dispatcher
	// go emmg.Dispatch("127.0.0.1:8888", emmChan)
	// for {
	// 	//we just send something for test
	// 	//in real work, you should assemble emm(entitle management message) section
	// 	emmChan <- []byte("Hello World")
	// }
	//channels := make([]utils.IdInterface, 3)

	channels := make([]*data.ProductChannel, 3)
	for i := 0; i < 3; i++ {
		channels[i] = new(data.ProductChannel)
		channels[i].Id = uint16(i + 1)
		channels[i].Name = "ch" + strconv.Itoa(i)
		channels[i].ServiceId = uint16(i)
		channels[i].TransportStreamId = uint16(i)
	}

	//var index uint16 = 1

	pos := utils.LowerBoundId[uint16](channels, 1)
	println(pos)

	ints := make([]int, 3)
	for i := 0; i < 3; i++ {
		ints[i] = i
	}
	p := utils.LowerBound[int](ints, 1)
	println(p)
	for i := 0; i < 3; i++ {
		println(channels[i].Id)
		println(channels[i].Name)
		println(ints[i])
	}
}
