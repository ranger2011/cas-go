package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	SectionTspktFlagSection = 0
	SectionTspktFlagTs      = 1
)

type Config struct {
	SuperCasId uint32
	DataMode   byte
}

var globalConfig Config

func GetSuperCasId() uint32 { return globalConfig.SuperCasId }
func GetDataMode() byte     { return globalConfig.DataMode }
func InitConfig() {
	globalConfig.SuperCasId = 0x1eb01168
	globalConfig.DataMode = SectionTspktFlagTs
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}
	json.Unmarshal(content, &globalConfig)
}
