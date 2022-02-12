package emmg

import "github.com/ranger2011/cas-go/comm"

func ChannelSetup(clientId uint32, channelId uint16, sectionTspktFlag byte, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmChannelSetup)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	return comm.AddParameterByte(buffer, EmmParaSectionTspktFlag, sectionTspktFlag)
}

func ChannelTest(clientId uint32, channelId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmChannelTest)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	return comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
}

func ChannelStatus(clientId uint32, channelId uint16, sectionTspktFlag byte, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmChannelStatus)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	return comm.AddParameterByte(buffer, EmmParaSectionTspktFlag, sectionTspktFlag)
}

func ChannelClose(clientId uint32, channelId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmChannelClose)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	return comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
}

func StreamSetup(clientId uint32, channelId uint16, streamId uint16, dataId uint16, dataType byte, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmStreamSetup)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
	comm.AddParameterWord(buffer, EmmParaDataId, dataId)
	return comm.AddParameterByte(buffer, EmmParaDataType, dataType)
}

func StreamTest(clientId uint32, channelId uint16, streamId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmStreamTest)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	return comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
}

func StreamStatus(clientId uint32, channelId uint16, streamId uint16, dataId uint16, dataType byte, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmStreamStatus)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
	comm.AddParameterWord(buffer, EmmParaDataId, dataId)
	return comm.AddParameterByte(buffer, EmmParaDataType, dataType)
}

func StreamCloseRequest(clientId uint32, channelId uint16, streamId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmStreamCloseRequest)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	return comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
}

func StreamCloseResponse(clientId uint32, channelId uint16, streamId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmStreamCloseResponse)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	return comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
}

func StreamError(clientId uint32, channelId uint16, streamId uint16, errorStatus uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmStreamError)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
	return comm.AddParameterWord(buffer, EmmParaErrorStatus, errorStatus)
}

func StreamBandwidthRequest(clientId uint32, channelId uint16, streamId uint16, bandwidth uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmStreamBwRequest)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
	return comm.AddParameterWord(buffer, EmmParaBandwidth, bandwidth)
}

func StreamBandwidthAllocation(clientId uint32, channelId uint16, streamId uint16, bandwidth uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmStreamBwAllocation)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
	return comm.AddParameterWord(buffer, EmmParaBandwidth, bandwidth)
}

func DataProvision(clientId uint32, channelId uint16, streamId uint16, dataId uint16, data []byte, dataLength uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, MessageEmmDataProvision)
	comm.AddParameterDWord(buffer, EmmParaClientId, clientId)
	comm.AddParameterWord(buffer, EmmParaDataChannelId, channelId)
	comm.AddParameterWord(buffer, EmmParaDataStreamId, streamId)
	comm.AddParameterWord(buffer, EmmParaDataId, dataId)
	return comm.AddParameterBlock(buffer, EmmParaDatagram, dataLength, data)
}

func GetClientId(message *comm.Message) (bool, uint32) {
	return message.GetParameterValueDWord(EmmParaClientId)
}

func GetSectionTsPktFlag(message *comm.Message) (bool, byte) {
	return message.GetParameterValueByte(EmmParaSectionTspktFlag)
}

func GetDataChannelId(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(EmmParaDataChannelId)
}

func GetDataStreamId(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(EmmParaDataStreamId)
}

func GetDataGram(message *comm.Message) (bool, []byte) {
	return message.GetParameterValueBlock(EmmParaDatagram)
}

func GetBandwidth(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(EmmParaBandwidth)
}

func GetDataType(message *comm.Message) (bool, byte) {
	return message.GetParameterValueByte(EmmParaDataType)
}

func GetDataId(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(EmmParaDataId)
}

func GetErrorStatus(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(EmmParaErrorStatus)
}

func GetErrorInformation(message *comm.Message) (bool, []byte) {
	return message.GetParameterValueBlock(EmmParaErrorInformation)
}
