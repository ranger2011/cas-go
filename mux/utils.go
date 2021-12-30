package mux

import (
	"net"
	"unsafe"

	"github.com/ranger2011/cas-go/mem"
)

func PackageUint16(buffer []byte, start int, value uint16) int {
	buffer[start] = byte(value >> 8)
	buffer[start+1] = byte(value)
	return 2
}

func ParseUint16(buffer []byte, start int) (value uint16, length int) {
	value = uint16(buffer[start])<<8 | uint16(buffer[start+1])
	length = 2
	return value, length
}

func PackageUint32(buffer []byte, start int, value uint32) int {
	buffer[start] = byte(value >> 24)
	buffer[start+1] = byte(value >> 16)
	buffer[start+2] = byte(value >> 8)
	buffer[start+3] = byte(value)
	return 4
}

func ParseUint32(buffer []byte, start int) (value uint32, length int) {
	value = uint32(buffer[0])<<24 | uint32(buffer[1])<<16 | uint32(buffer[2])<<8 | uint32(buffer[3])
	length = 4
	return value, length
}

func PackageUint64(buffer []byte, start int, value uint64) int {
	buffer[start] = byte(value >> 56)
	buffer[start+1] = byte(value >> 48)
	buffer[start+2] = byte(value >> 40)
	buffer[start+3] = byte(value >> 32)
	buffer[start+4] = byte(value >> 24)
	buffer[start+5] = byte(value >> 16)
	buffer[start+6] = byte(value >> 8)
	buffer[start+7] = byte(value)
	return 8
}

func ParseUint64(buffer []byte, start int) (value uint64, length int) {
	value = uint64(buffer[0])<<56 | uint64(buffer[1])<<48 | uint64(buffer[2])<<40 | uint64(buffer[3])<<32 |
		uint64(buffer[4])<<24 | uint64(buffer[5])<<16 | uint64(buffer[6])<<8 | uint64(buffer[7])
	length = 8
	return value, length
}

func MemCpy(dest, src []byte, destStart, srcStart int, len int) int {
	mem.Memcpy(unsafe.Pointer(&dest[destStart]), unsafe.Pointer(&src[srcStart]), len)
	return len
}

func SendTcp(buffer []byte, length int, conn net.Conn) bool {
	sendBuffer := make([]byte, length)
	MemCpy(sendBuffer, buffer, 0, 0, length)
	_, err := conn.Write(sendBuffer)
	return err == nil
}
