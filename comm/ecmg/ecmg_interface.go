package ecmg

import (
	"github.com/ranger2011/cas-go/comm"
	"github.com/ranger2011/cas-go/config"
	"github.com/ranger2011/cas-go/utils"
)

func ChannelSetup(channelId uint16, superCasId uint32, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_SETUP)

	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return comm.AddParameterDWord(buffer, ECM_PARA_SUPER_CAS_ID, superCasId)
}

func ChannelTest(channelId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_TEST)

	return comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
}

func ChannelStatus(channelId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_STATUS)

	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	comm.AddParameterByte(buffer, ECM_PARA_SECTION_TSPKT_FLAG, config.GetDataMode())
	comm.AddParameterWord(buffer, ECM_PARA_AC_DELAY_START, AcDelayStart)
	comm.AddParameterWord(buffer, ECM_PARA_AC_DELAY_STOP, AcDelayStop)
	comm.AddParameterWord(buffer, ECM_PARA_DELAY_START, DelayStart)
	comm.AddParameterWord(buffer, ECM_PARA_DELAY_STOP, DelayStop)
	comm.AddParameterWord(buffer, ECM_PARA_TRANSITION_DELAY_START, TransitionDelayStart)
	comm.AddParameterWord(buffer, ECM_PARA_TRANSITION_DELAY_STOP, TransitionDelayStop)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_REP_PERIOD, EcmRepPeriod)
	comm.AddParameterWord(buffer, ECM_PARA_MAX_STREAMS, MaxStreams)
	comm.AddParameterWord(buffer, ECM_PARA_MIN_CP_DURATION, MinCpDuration)
	comm.AddParameterByte(buffer, ECM_PARA_LEAD_CW, LEAD_CW)
	comm.AddParameterByte(buffer, ECM_PARA_CW_PER_MSG, CW_PER_MSG)
	return comm.AddParameterWord(buffer, ECM_PARA_MAX_COMP_TIME, MaxCompTime)
}

func ChannelClose(channelId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_CLOSE)

	return comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
}

func ChannelError(channelId uint16, errorStatus uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CHANNEL_ERROR)

	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return comm.AddParameterWord(buffer, ECM_PARA_ERROR_STATUS, errorStatus)
}

func StreamSetup(channelId uint16, streamId uint16, ecmId uint16, nominalCpDuration uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_SETUP)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	return comm.AddParameterWord(buffer, ECM_PARA_NOMINAL_CP_DURATION, nominalCpDuration)
}

func StreamTest(channelId uint16, streamId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_TEST)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return comm.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
}

func StreamStatus(channelId uint16, streamId uint16, ecmId uint16, accessCriteriaMode byte, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_STATUS)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	return comm.AddParameterByte(buffer, ECM_PARA_ACCESS_CRITERIA_TRANSFER_MODE, accessCriteriaMode)
}

func StreamCloseRequest(channelId uint16, streamId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_CLOSE_REQUEST)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return comm.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
}

func StreamCloseResponse(channelId uint16, streamId uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_CLOSE_RESPONSE)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	return comm.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
}

func StreamError(channelId uint16, streamId uint16, errorStatus uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_STREAM_ERROR)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	return comm.AddParameterWord(buffer, ECM_PARA_ERROR_STATUS, errorStatus)
}

func CwProvition(channelId uint16, streamId uint16, cpNumber uint16,
	cpCwCombination1 []byte, cpCwCombination2 []byte, cpCwLength uint16, cpDuration uint16,
	accessCriteria []byte, accessCriteriaLength uint16, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_CW_PROVISION)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	comm.AddParameterBlock(buffer, ECM_PARA_CP_CW_COMBINATION, cpCwLength, cpCwCombination1)
	comm.AddParameterBlock(buffer, ECM_PARA_CP_CW_COMBINATION, cpCwLength, cpCwCombination2)
	comm.AddParameterWord(buffer, ECM_PARA_CP_DURATION, cpDuration)
	return comm.AddParameterBlock(buffer, ECM_PARA_ACCESS_CRITERIA, accessCriteriaLength, accessCriteria)
}

func EcmResponse(channelId uint16, streamId uint16, cpNumber uint16, ecmLen uint16, ecmData []byte, buffer []byte) int {
	comm.PackageMessageHead(buffer, ECM_MESSAGE_TYPE_ECM_RESPONSE)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_CHANNEL_ID, channelId)
	comm.AddParameterWord(buffer, ECM_PARA_ECM_STREAM_ID, streamId)
	comm.AddParameterWord(buffer, ECM_PARA_CP_NUMBER, cpNumber)
	return comm.AddParameterBlock(buffer, ECM_PARA_ECM_DATAGRAM, ecmLen, ecmData)
}

func GetSuperCasId(message *comm.Message) (bool, uint32) {
	return message.GetParameterValueDWord(ECM_PARA_SUPER_CAS_ID)
}

func GetSectionTsPktFlag(message *comm.Message) (bool, byte) {
	return message.GetParameterValueByte(ECM_PARA_SECTION_TSPKT_FLAG)
}

func GetDelayStart(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_DELAY_START)
}

func GetDelayStop(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_DELAY_STOP)
}

func GetTransitionDelayStart(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_TRANSITION_DELAY_START)
}

func GetTransitionDelayStop(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_TRANSITION_DELAY_STOP)
}

func GetEcmRepPeriod(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ECM_REP_PERIOD)
}

func GetMaxStreams(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_MAX_STREAMS)
}

func GetMinCpDuration(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_MIN_CP_DURATION)
}

func GetLeadCw(message *comm.Message) (bool, byte) {
	return message.GetParameterValueByte(ECM_PARA_LEAD_CW)
}

func GetCwPerMsg(message *comm.Message) (bool, byte) {
	return message.GetParameterValueByte(ECM_PARA_CW_PER_MSG)
}

func GetMaxCompTime(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_MAX_COMP_TIME)
}

func GetAccessCriteria(message *comm.Message) (bool, []byte) {
	return message.GetParameterValueBlock(ECM_PARA_ACCESS_CRITERIA)
}

func GetDataChannelId(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ECM_CHANNEL_ID)
}

func GetDataStreamId(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ECM_STREAM_ID)
}

func GetNominalCpDuration(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_NOMINAL_CP_DURATION)
}

func GetAccessCriteriaTransferMode(message *comm.Message) (bool, byte) {
	return message.GetParameterValueByte(ECM_PARA_ACCESS_CRITERIA_TRANSFER_MODE)
}

func GetCpNumber(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_CP_NUMBER)
}

func GetCpDuration(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_CP_DURATION)
}

func GetCw(message *comm.Message) (bool, []byte) {
	//cp cw combination: 0, 1 byte for cp number, 2~8 for cw
	//I don't use cp number, so get cw only
	found, result := message.GetParameterValueBlock(ECM_PARA_CP_CW_COMBINATION)
	if !found {
		return false, nil
	}
	var length = len(result) / 10 * 8
	var cw = make([]byte, length)
	var cp, _ = utils.ParseUint16(result, 0)
	if cp%2 == 0 {
		utils.MemCpy(cw, result, 0, 12, 8)
		if length > 8 {
			utils.MemCpy(cw, result, 0, 2, 8)
		}
	} else {
		utils.MemCpy(cw, result, 0, 2, 8)
		if length > 8 {
			utils.MemCpy(cw, result, 8, 12, 8)
		}
	}

	return true, cw
}

func GetDataGram(message *comm.Message) (bool, []byte) {
	return message.GetParameterValueBlock(ECM_PARA_ECM_DATAGRAM)
}

func GetAcDelayStart(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_AC_DELAY_START)
}

func GetAcDelayStop(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_AC_DELAY_STOP)
}

func GetEcmId(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ECM_ID)
}

func GetErrorStatus(message *comm.Message) (bool, uint16) {
	return message.GetParameterValueWord(ECM_PARA_ERROR_STATUS)
}

func GetErrorInformation(message *comm.Message) (bool, []byte) {
	return message.GetParameterValueBlock(ECM_PARA_ERROR_INFORMATION)
}
