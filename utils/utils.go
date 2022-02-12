package utils

import (
	"crypto/aes"
	"fmt"
	"net"
	"os"
	"reflect"
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
	value = uint32(buffer[start])<<24 | uint32(buffer[start+1])<<16 | uint32(buffer[start+2])<<8 | uint32(buffer[start+3])
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

func MemCpy(dst, src []byte, destStart, srcStart int, len int) int {
	mem.Memcpy(unsafe.Pointer(&dst[destStart]), unsafe.Pointer(&src[srcStart]), len)
	return len
}

func MemComp(dst, src []byte) bool {
	if len(dst) != len(src) {
		return false
	}
	return mem.Memcmp(unsafe.Pointer(&dst[0]), unsafe.Pointer(&src[0]), len(dst)) == 0
}

func SendTcp(buffer []byte, length int, conn net.Conn) bool {
	sendBuffer := make([]byte, length)
	MemCpy(sendBuffer, buffer, 0, 0, length)
	_, err := conn.Write(sendBuffer)
	return err == nil
}

func DecryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func EncryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	encrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Encrypt(encrypted[bs:be], data[bs:be])
	}

	return encrypted
}

func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func PrintHex(byteArray []byte) {
	const lineSize int = 16
	arrayLen := len(byteArray)
	lineNumber := arrayLen / lineSize
	if arrayLen%lineSize > 0 {
		lineNumber++
	}
	for i := 0; i < lineNumber; i++ {
		len := lineSize
		if i == lineNumber-1 {
			len = arrayLen - lineSize*i
		}
		for j := 0; j < len; j++ {
			fmt.Printf("  %02x", byteArray[i*16+j])
		}
		fmt.Printf("\n")
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
