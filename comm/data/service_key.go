package data

import (
	"os"

	"github.com/ranger2011/cas-go/utils"
)

type ServiceKey struct {
	Odd     bool
	OddKey  []byte
	EvenKey []byte
}

var globalServiceKey ServiceKey

func LoadServiceKey() {
	f, err := os.Open("gsService.dat")
	if err != nil {
		return
	}
	defer f.Close()
	buffer := make([]byte, 64)
	n, err := f.Read(buffer)
	if err != nil {
		return
	}
	if n != 33 {
		return
	}
	if buffer[0] == 1 {
		globalServiceKey.Odd = true
	} else {
		globalServiceKey.Odd = false
	}
	encryptKey := []byte{11, 22, 33, 44, 55, 66, 77, 88, 99, 12, 23, 34, 45, 56, 67, 78}
	encrypted := make([]byte, 16)
	utils.MemCpy(encrypted, buffer, 0, 1, 16)
	globalServiceKey.OddKey = utils.DecryptAes128Ecb(encrypted, encryptKey)
	utils.MemCpy(encrypted, buffer, 0, 17, 16)
	globalServiceKey.EvenKey = utils.DecryptAes128Ecb(encrypted, encryptKey)
}

func IsCurrentKeyOdd() bool {
	return globalServiceKey.Odd
}

func GetCurrentKey() []byte {
	if globalServiceKey.Odd {
		return globalServiceKey.OddKey
	} else {
		return globalServiceKey.EvenKey
	}
}
