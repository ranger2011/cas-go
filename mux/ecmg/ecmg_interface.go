package ecmg

import "github.com/ranger2011/cas-go/mux"

func ChannelSetup(channelId uint16, superCasId uint32, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_SETUP)

	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return mux.AddParameterDWord(buffer, ECM_PARA_SUPER_CAS_ID, superCasId)
}

func ChannelTest(channelId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_TEST)

	return mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
}

func ChannelStatus(channelId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_STATUS)

	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	mux.AddParameterByte(buffer, ECM_PARA_SECTION_TSPKT_FLAG, mux.SectionTspktFlagSection)
	mux.AddParameterWord(buffer, ECM_PARA_AC_DELAY_START, AcDelayStart)
	mux.AddParameterWord(buffer, ECM_PARA_AC_DELAY_STOP, AcDelayStop)
	mux.AddParameterWord(buffer, ECM_PARA_DELAY_START, DelayStart)
	mux.AddParameterWord(buffer, ECM_PARA_DELAY_STOP, DelayStop)
	mux.AddParameterWord(buffer, ECM_PARA_TRANSITION_DELAY_START, TransitionDelayStart)
	mux.AddParameterWord(buffer, ECM_PARA_TRANSITION_DELAY_STOP, TransitionDelayStop)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_REP_PERIOD, EcmRepPeriod)
	mux.AddParameterWord(buffer, ECM_PARA_MAX_STREAMS, MaxStreams)
	mux.AddParameterWord(buffer, ECM_PARA_MIN_CP_DURATION, MinCpDuration)
	mux.AddParameterByte(buffer, ECM_PARA_LEAD_CW, LEAD_CW)
	mux.AddParameterByte(buffer, ECM_PARA_CW_PER_MSG, CW_PER_MSG)
	return mux.AddParameterWord(buffer, ECM_PARA_MAX_COMP_TIME, MaxCompTime)
}

func ChannelClose(channelId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_CLOSE)

	return mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
}

func ChannelError(channelId uint16, errorStatus uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_ERROR)

	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return mux.AddParameterWord(buffer, ECM_PARA_ERROR_STATUS, errorStatus)
}

func StreamSetup(channelId uint16, streamId uint16, ecmId uint16, nominalCpDuration uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_SETUP)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	return mux.AddParameterWord(buffer, ECM_PARA_NOMINAL_CP_DURATION, nominalCpDuration)
}

func StreamTest(channelId uint16, streamId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_TEST)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return mux.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
}

func StreamStatus(channelId uint16, streamId uint16, ecmId uint16, accessCriteriaMode byte, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_STATUS)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	return mux.AddParameterByte(buffer, ECM_PARA_ACCESS_CRITERIA_TRANSFER_MODE, accessCriteriaMode)
}

func StreamCloseRequest(channelId uint16, streamId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_CLOSE_REQUEST)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return mux.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
}

func StreamCloseResponse(channelId uint16, streamId uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_CLOSE_RESPONSE)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return mux.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
}

func StreamError(channelId uint16, streamId uint16, errorStatus uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_ERROR)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	return mux.AddParameterWord(buffer, ECM_PARA_ERROR_STATUS, errorStatus)
}

func CwProvition(channelId uint16, streamId uint16, cpNumber uint16,
	cpCwCombination1 []byte, cpCwCombination2 []byte, cpCwLength uint16, cpDuration uint16,
	accessCriteria []byte, accessCriteriaLength uint16, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CW_PROVISION)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	mux.AddParameterBlock(buffer, ECM_PARA_CP_CW_COMBINATION, cpCwLength, cpCwCombination1)
	mux.AddParameterBlock(buffer, ECM_PARA_CP_CW_COMBINATION, cpCwLength, cpCwCombination2)
	mux.AddParameterWord(buffer, ECM_PARA_CP_DURATION, cpDuration)
	return mux.AddParameterBlock(buffer, ECM_PARA_ACCESS_CRITERIA, accessCriteriaLength, accessCriteria)
}

func EcmResponse(channelId uint16, streamId uint16, cpNumber uint16, ecmData []byte, buffer []byte) int {
	mux.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_ECM_RESPONSE)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	mux.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	mux.AddParameterWord(buffer, ECM_PARA_CP_NUMBER, cpNumber)
	return mux.AddParameterBlock(buffer, ECM_PARA_ECM_DATAGRAM, uint16(len(ecmData)), ecmData)
}

func GetSuperCasId(message *mux.Message) (bool, uint32) {
	return message.GetParameterValueDWord(ECM_PARA_SUPER_CAS_ID)
}

func GetSectionTsPktFlag(message *mux.Message) (bool, byte) {
	return message.GetParameterValueByte(ECM_PARA_SECTION_TSPKT_FLAG)
}

func GetDelayStart(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_DELAY_START)
}

func GetDelayStop(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_DELAY_STOP)
}

func GetTransitionDelayStart(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_TRANSITION_DELAY_START)
}

func GetTransitionDelayStop(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_TRANSITION_DELAY_STOP)
}

func GetEcmRepPeriod(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ECM_REP_PERIOD)
}

func GetMaxStreams(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_MAX_STREAMS)
}

func GetMinCpDuration(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_MIN_CP_DURATION)
}

func GetLeadCw(message *mux.Message) (bool, byte) {
	return message.GetParameterValueByte(ECM_PARA_LEAD_CW)
}

func GetCwPerMsg(message *mux.Message) (bool, byte) {
	return message.GetParameterValueByte(ECM_PARA_CW_PER_MSG)
}

func GetMaxCompTime(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_MAX_COMP_TIME)
}

func GetAccessCriteria(message *mux.Message) (bool, []byte) {
	return message.GetParameterValueBlock(ECM_PARA_ACCESS_CRITERIA)
}

func GetDataChannelId(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ECM_CHANNEL_ID)
}

func GetDataStreamId(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ECM_STREAM_ID)
}

func GetNominalCpDuration(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_NOMINAL_CP_DURATION)
}

func GetAccessCriteriaTransferMode(message *mux.Message) (bool, byte) {
	return message.GetParameterValueByte(ECM_PARA_ACCESS_CRITERIA_TRANSFER_MODE)
}

func GetCpNumber(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_CP_NUMBER)
}

func GetCpDuration(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_CP_DURATION)
}

func GetCw(message *mux.Message) (bool, []byte) {
	//cp cw combination: 0, 1 byte for cp number, 2~8 for cw
	//i don't use cp number, so get cw only
	found, result := message.GetParameterValueBlock(ECM_PARA_CP_CW_COMBINATION)
	if !found {
		return false, nil
	}
	var length = len(result) / 10 * 8
	var cw = make([]byte, length)
	mux.MemCpy(cw, result, 0, 2, 8)
	if length > 8 {
		mux.MemCpy(cw, result, 8, 12, 8)
	}
	return true, cw
}

func GetDataGram(message *mux.Message) (bool, []byte) {
	return message.GetParameterValueBlock(ECM_PARA_ECM_DATAGRAM)
}

func GetAcDelayStart(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_AC_DELAY_START)
}

func GetAcDelayStop(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_AC_DELAY_STOP)
}

func GetEcmId(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ECM_ID)
}

func GetErrorStatus(message *mux.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ERROR_STATUS)
}

func GetErrorInformation(message *mux.Message) (bool, []byte) {
	return message.GetParameterValueBlock(ECM_PARA_ERROR_INFORMATION)
}
