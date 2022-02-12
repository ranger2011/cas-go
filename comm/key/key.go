package key

import (
	"crypto/md5"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/ranger2011/cas-go/utils"
)

const KeyLength uint8 = 16
const GroupSizeVersion8 uint8 = 25
const GroupSizeVersion9 uint8 = 100

type Key struct {
	Id      uint32
	Content []byte
}

type KeyCollection struct {
	Type       string
	Collection []*Key
}

type KeyCollectionGroup struct {
	Label          string
	Version        uint8
	KeyCollections []*KeyCollection
}

var KeyCollectionGroupArray []*KeyCollectionGroup

func constructDirName(version uint8) string {
	return "version" + strconv.Itoa(int(version))
}

func constructFileNames(version uint8, keyType string, label string) (string, string) {
	keyFileName := constructDirName(version) + "/" + label + "_" + keyType + ".dat"
	md5FileName := constructDirName(version) + "/" + label + "_" + keyType + ".md5"
	return keyFileName, md5FileName
}

func CheckLabel(version uint8) (bool, string) {
	dirname := constructDirName(version)
	if !utils.PathExists(dirname) {
		return false, ""
	}
	items, _ := ioutil.ReadDir(dirname)
	if len(items) == 0 {
		return false, ""
	}
	for _, item := range items {
		if !item.IsDir() {
			tags := strings.Split(item.Name(), "_")
			label := tags[0]
			return true, label
		}
	}
	return false, ""
}

func parseKey(bytes []byte) *Key {
	result := new(Key)
	result.Id, _ = utils.ParseUint32(bytes, 0)
	utils.MemCpy(result.Content, bytes, 0, 4, int(KeyLength))
	return result
}

func LoadKeyData(version uint8, keyType string, label string) (*KeyCollection, bool) {
	keyFileName, md5FileName := constructFileNames(version, keyType, label)

	keyFile, err := os.Open(keyFileName)
	if err != nil {
		return nil, false
	}
	defer keyFile.Close()

	md5File, err := os.Open(md5FileName)
	if err != nil {
		return nil, false
	}
	defer md5File.Close()
	md5Buffer := make([]byte, KeyLength)
	n, err := md5File.Read(md5Buffer)
	if err != nil || uint8(n) != KeyLength {
		return nil, false
	}

	result := new(KeyCollection)
	result.Type = keyType

	buffer := make([]byte, 20)
	cache := buffer[:4]
	keyFile.ReadAt(cache, 4)
	keyNumber, _ := utils.ParseUint32(cache, 0)
	result.Collection = make([]*Key, keyNumber)
	digest := md5.New()
	var byteCount int64 = 4
	for i := 0; i < int(keyNumber); i++ {
		n, err := keyFile.ReadAt(buffer, byteCount)
		if err != nil {
			return nil, false
		}
		if n != 20 {
			return nil, false
		}
		digest.Write(buffer)
		result.Collection[i] = parseKey(buffer)
	}
	md5Result := digest.Sum(nil)
	if !utils.MemComp(md5Result, md5Buffer) {
		return nil, false
	}
	return result, true
}

func LoadKeyDataGroup(version uint8) *KeyCollectionGroup {
	found, label := CheckLabel(version)
	if !found {
		return nil
	}
	ini, ok := LoadKeyData(version, "ini", label)
	if !ok {
		return nil
	}
	lic, ok := LoadKeyData(version, "lic", label)
	if !ok {
		return nil
	}
	psk, ok := LoadKeyData(version, "psk", label)
	if !ok {
		return nil
	}
	result := new(KeyCollectionGroup)
	result.Label = label
	result.Version = version
	if version == 6 {
		result.KeyCollections = make([]*KeyCollection, 3)
	} else {
		result.KeyCollections = make([]*KeyCollection, 4)
	}
	result.KeyCollections[0] = ini
	result.KeyCollections[1] = lic
	result.KeyCollections[2] = psk
	if version == 6 {
		return result
	}

	gos, ok := LoadKeyData(version, "gos", label)
	if !ok {
		return nil
	}
	result.KeyCollections[3] = gos
	return result
}

func InitKeyLib() bool {
	version6 := LoadKeyDataGroup(6)
	version7 := LoadKeyDataGroup(7)
	version8 := LoadKeyDataGroup(8)
	version9 := LoadKeyDataGroup(9)
	count := 0
	if version6 != nil {
		count++
	}
	if version7 != nil {
		count++
	}
	if version8 != nil {
		count++
	}
	if version9 != nil {
		count++
	}
	if count == 0 {
		return false
	}
	KeyCollectionGroupArray = make([]*KeyCollectionGroup, count)
	count = 0
	if version6 != nil {
		KeyCollectionGroupArray[count] = version6
		count++
	}
	if version7 != nil {
		KeyCollectionGroupArray[count] = version7
		count++
	}
	if version8 != nil {
		KeyCollectionGroupArray[count] = version8
		count++
	}
	if version9 != nil {
		KeyCollectionGroupArray[count] = version9
	}
	return true
}
