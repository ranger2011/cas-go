package emmg

import "github.com/ranger2011/cas-go/mux"

func ChannelSetup(clientId uint32, channelId uint16, sectionTspktFlag byte, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_CHANNEL_SETUP)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	return mux.AddParameterByte(buffer, EMM_PARA_SECTION_TSPKT_FLAG, sectionTspktFlag)
}

func ChannelTest(clientId uint32, channelId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_CHANNEL_TEST)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	return mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
}

func ChannelStatus(clientId uint32, channelId uint16, sectionTspktFlag byte, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_CHANNEL_STATUS)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	return mux.AddParameterByte(buffer, EMM_PARA_SECTION_TSPKT_FLAG, sectionTspktFlag)
}

func ChannelClose(clientId uint32, channelId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_CHANNEL_CLOSE)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	return mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
}

func StreamSetup(clientId uint32, channelId uint16, streamId uint16, dataId uint16, dataType byte, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_STREAM_SETUP)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_ID, dataId)
	return mux.AddParameterByte(buffer, EMM_PARA_DATA_TYPE, dataType)
}

func StreamTest(clientId uint32, channelId uint16, streamId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_STREAM_TEST)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	return mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
}

func StreamStatus(clientId uint32, channelId uint16, streamId uint16, dataId uint16, dataType byte, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_STREAM_STATUS)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_ID, dataId)
	return mux.AddParameterByte(buffer, EMM_PARA_DATA_TYPE, dataType)
}

func StreamCloseRequest(clientId uint32, channelId uint16, streamId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_STREAM_CLOSE_REQUEST)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	return mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
}

func StreamCloseResponse(clientId uint32, channelId uint16, streamId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_STREAM_CLOSE_RESPONSE)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	return mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
}

func StreamError(clientId uint32, channelId uint16, streamId uint16, errorStatus uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_STREAM_ERROR)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
	return mux.AddParameterWord(buffer, EMM_PARA_ERROR_STATUS, errorStatus)
}

func StreamBandwidthRequest(clientId uint32, channelId uint16, streamId uint16, bandwidth uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_STREAM_BW_REQUEST)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
	return mux.AddParameterWord(buffer, EMM_PARA_BANDWIDTH, bandwidth)
}

func StreamBandwidthAllocation(clientId uint32, channelId uint16, streamId uint16, bandwidth uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_STREAM_BW_ALLOCATION)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
	return mux.AddParameterWord(buffer, EMM_PARA_BANDWIDTH, bandwidth)
}

func DataProvision(clientId uint32, channelId uint16, streamId uint16, dataId uint16, data []byte, dataLength uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, MESSAGE_EMM_DATA_PROVISION)
	mux.AddParameterDWord(buffer, EMM_PARA_CLIENT_ID, clientId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_STREAM_ID, streamId)
	mux.AddParameterWord(buffer, EMM_PARA_DATA_ID, dataId)
	return mux.AddParameterBlock(buffer, EMM_PARA_DATAGRAM, dataLength, data)
}

func GetClientId(message *mux.Message) (bool, uint32) {
	return message.GetParameterValueDWord(EMM_PARA_CLIENT_ID)
}

func GetSectionTsPktFlag(message *mux.Message) (bool, byte) {
	return message.GetParameterValueByte(EMM_PARA_SECTION_TSPKT_FLAG)
}

func GetDataChannelId(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(EMM_PARA_DATA_CHANNEL_ID)
}

func GetDataStreamId(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(EMM_PARA_DATA_STREAM_ID)
}

func GetDataGram(message *mux.Message) (bool, []byte) {
	return message.GetParameterValueBlock(EMM_PARA_DATAGRAM)
}

func GetBandwidth(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(EMM_PARA_BANDWIDTH)
}

func GetDataType(message *mux.Message) (bool, byte) {
	return message.GetParameterValueByte(EMM_PARA_DATA_TYPE)
}

func GetDataId(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(EMM_PARA_DATA_ID)
}

func GetErrorStatus(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(EMM_PARA_ERROR_STATUS)
}

func GetErrorInformation(message *mux.Message) (bool, []byte) {
	return message.GetParameterValueBlock(EMM_PARA_ERROR_INFORMATION)
}
