package main

import (
	"fmt"

	"github.com/ranger2011/cas-go/mux"
	"github.com/ranger2011/cas-go/mux/ecmg"
	"github.com/ranger2011/cas-go/mux/emmg"
)

func main() {
	fmt.Println("Conditional access system demo....")
	fmt.Printf("The protocol verison is %d.\n", mux.ProtocolVersion)
	//start a ecmg server
	go ecmg.StartEcmg(5555)
	emmChan := make(chan []byte)
	//start a emmg dispatcher
	go emmg.Dispatch("127.0.0.1:8888", emmChan)
	for {
		//we just send something for test
		//in real work, you should assemble emm(entitle management message) section
		emmChan <- []byte("Hello World")
	}
}
