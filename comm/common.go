package comm

import "github.com/ranger2011/cas-go/utils"

const (
	ProtocolVersion   = 2
	MessageHeadLength = 5
	ParameterLength   = 4

	SuperCasId = 0x12340001
)

type MessageHeader struct {
	Version byte
	Type    uint16
	Length  uint16
}

type Parameter struct {
	Type    uint16
	Length  uint16
	Content []byte
}

type Message struct {
	Header     MessageHeader
	Parameters []Parameter
}

func PackageMessageHead(buffer []byte, messageType uint16) int {
	buffer[0] = ProtocolVersion
	utils.PackageUint16(buffer, 1, messageType)
	buffer[3] = 0
	buffer[4] = 0
	return MessageHeadLength
}

func (header *MessageHeader) ParseMessageHead(buffer []byte) int {
	header.Version = buffer[0]
	header.Type, _ = utils.ParseUint16(buffer, 1)
	header.Length, _ = utils.ParseUint16(buffer, 3)
	return MessageHeadLength
}

func addParameterContent(buffer []byte, parameter *Parameter) int {
	var length int = int(buffer[3])<<8 + int(buffer[4]) + MessageHeadLength

	length += utils.PackageUint16(buffer, length, parameter.Type)
	length += utils.PackageUint16(buffer, length, parameter.Length)
	length += utils.MemCpy(buffer, parameter.Content, length, 0, int(parameter.Length))

	length -= MessageHeadLength
	utils.PackageUint16(buffer, 3, uint16(length))
	return length + MessageHeadLength
}

func AddParameterByte(buffer []byte, paraType uint16, value byte) int {
	var para Parameter
	para.Type = paraType
	para.Length = 1
	para.Content = make([]byte, 1)
	para.Content[0] = value
	return addParameterContent(buffer, &para)
}

func AddParameterWord(buffer []byte, paraType uint16, value uint16) int {
	var para Parameter
	para.Type = paraType
	para.Length = 2
	para.Content = make([]byte, 2)
	utils.PackageUint16(para.Content, 0, value)
	return addParameterContent(buffer, &para)
}

func AddParameterDWord(buffer []byte, paraType uint16, value uint32) int {
	var para Parameter
	para.Type = paraType
	para.Length = 4
	para.Content = make([]byte, 4)
	utils.PackageUint32(para.Content, 0, value)
	return addParameterContent(buffer, &para)
}

func AddParameterQWord(buffer []byte, paraType uint16, value uint64) int {
	var para Parameter
	para.Type = paraType
	para.Length = 8
	para.Content = make([]byte, 8)
	utils.PackageUint64(para.Content, 0, value)
	return addParameterContent(buffer, &para)
}

func AddParameterBlock(buffer []byte, paraType uint16, paraLength uint16, block []byte) int {
	var para Parameter
	para.Type = paraType
	para.Length = paraLength
	para.Content = make([]byte, paraLength)
	utils.MemCpy(para.Content, block, 0, 0, int(paraLength))
	return addParameterContent(buffer, &para)
}

func (parameter *Parameter) ParseParameterContent(buffer []byte, start int) int {
	var length, count int

	parameter.Type, count = utils.ParseUint16(buffer, start+length)
	length += count
	parameter.Length, count = utils.ParseUint16(buffer, start+length)
	length += count
	parameter.Content = make([]byte, parameter.Length)
	length += utils.MemCpy(parameter.Content, buffer, 0, start+length, int(parameter.Length))
	return length
}

func (message *Message) Package(buffer []byte) int {
	var length int
	length += PackageMessageHead(buffer, message.Header.Type)
	for i := 0; i < len(message.Parameters); i++ {
		length += addParameterContent(buffer, &message.Parameters[i])
	}
	return length
}

func (message *Message) Parse(buffer []byte, receivedLen int) bool {
	var length int
	length += message.Header.ParseMessageHead(buffer)
	if message.Header.Length+MessageHeadLength != uint16(receivedLen) {
		return false
	}
	message.Parameters = nil
	for {
		if length-MessageHeadLength >= int(message.Header.Length) {
			break
		}
		var parameter Parameter
		length += parameter.ParseParameterContent(buffer, length)
		message.Parameters = append(message.Parameters, parameter)
	}
	return true
}

func (message *Message) findParameter(parameterType uint16) []Parameter {
	var result []Parameter
	for i := 0; i < len(message.Parameters); i++ {
		if message.Parameters[i].Type == parameterType {
			result = append(result, message.Parameters[i])
		}
	}
	return result
}

func (message *Message) GetParameterValueByte(parameterType uint16) (bool, byte) {
	parameters := message.findParameter(parameterType)
	if len(parameters) == 0 {
		return false, 0
	}
	return true, parameters[0].Content[0]
}

func (message *Message) GetParameterValueWord(parameterType uint16) (bool, uint16) {
	parameters := message.findParameter(parameterType)
	if len(parameters) == 0 {
		return false, 0
	}
	value, _ := utils.ParseUint16(parameters[0].Content, 0)
	return true, value
}

func (message *Message) GetParameterValueDWord(parameterType uint16) (bool, uint32) {
	parameters := message.findParameter(parameterType)
	if len(parameters) == 0 {
		return false, 0
	}
	value, _ := utils.ParseUint32(parameters[0].Content, 0)
	return true, value
}

func (message *Message) GetParameterValueQWord(parameterType uint16) (bool, uint64) {
	parameters := message.findParameter(parameterType)
	if len(parameters) == 0 {
		return false, 0
	}
	value, _ := utils.ParseUint64(parameters[0].Content, 0)
	return true, value
}

func (message *Message) GetParameterValueBlock(parameterType uint16) (bool, []byte) {
	parameters := message.findParameter(parameterType)
	if len(parameters) == 0 {
		return false, nil
	}
	var length int = 0
	for i := 0; i < len(parameters); i++ {
		length += len(parameters[i].Content)
	}
	result := make([]byte, length)
	var count = 0
	for i := 0; i < len(parameters); i++ {
		utils.MemCpy(result, parameters[i].Content, count, 0, len(parameters[i].Content))
		count += len(parameters[i].Content)
	}
	return true, result
}
