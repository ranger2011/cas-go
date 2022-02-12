package emmg

const (
	//channel and stream constants

	MessageEmmChannelSetup        uint16 = 0x0011
	MessageEmmChannelTest         uint16 = 0x0012
	MessageEmmChannelStatus       uint16 = 0x0013
	MessageEmmChannelClose        uint16 = 0x0014
	MessageEmmChannelError        uint16 = 0x0015
	MessageEmmStreamSetup         uint16 = 0x0111
	MessageEmmStreamTest          uint16 = 0x0112
	MessageEmmStreamStatus        uint16 = 0x0113
	MessageEmmStreamCloseRequest  uint16 = 0x0114
	MessageEmmStreamCloseResponse uint16 = 0x0115
	MessageEmmStreamError         uint16 = 0x0116
	MessageEmmStreamBwRequest     uint16 = 0x0117
	MessageEmmStreamBwAllocation  uint16 = 0x0118
	MessageEmmDataProvision       uint16 = 0x0211

	//error status

	ErrorInvalidMessagePackage             uint16 = 0x0001
	ErrorUnsupportedProtocolVersion        uint16 = 0x0002
	ErrorUnknownMessageTypeValue           uint16 = 0x0003
	ErrorMessageTooLong                    uint16 = 0x0004
	ErrorUnknownDataStreamIdValue          uint16 = 0x0005
	ErrorUnknownDataChannelIdValue         uint16 = 0x0006
	ErrorTooManyChannelsOnThisMux          uint16 = 0x0007
	ErrorTooManyDataStreamsOnThisChannel   uint16 = 0x0008
	ErrorTooManyDataStreamsOnThisMux       uint16 = 0x0009
	ErrorUnknownParameterType              uint16 = 0x000a
	ErrorInconsistentLengthForDvbParameter uint16 = 0x000b
	ErrorMissingMandatoryDvbParameter      uint16 = 0x000c
	ErrorInvalidValueForDvbParameter       uint16 = 0x000d
	ErrorUnknownClientIdValue              uint16 = 0x000e
	ErrorExceededBandwidth                 uint16 = 0x000f
	ErrorUnknownDataIdValue                uint16 = 0x0010
	ErrorDataChannelIdValueAlreadyInUse    uint16 = 0x0011
	ErrorDataStreamIdValueAlreadyInUse     uint16 = 0x0012
	ErrorDataIdValueAlreadyInuse           uint16 = 0x0013
	ErrorClientIdValueAlreadyInUse         uint16 = 0x0014
	ErrorUnknownError                      uint16 = 0x7000
	ErrorUnrecoverableError                uint16 = 0x7001

	//parameters

	EmmParaClientId         uint16 = 0x0001
	EmmParaSectionTspktFlag uint16 = 0x0002
	EmmParaDataChannelId    uint16 = 0x0003
	EmmParaDataStreamId     uint16 = 0x0004
	EmmParaDatagram         uint16 = 0x0005
	EmmParaBandwidth        uint16 = 0x0006
	EmmParaDataType         uint16 = 0x0007
	EmmParaDataId           uint16 = 0x0008
	EmmParaErrorStatus      uint16 = 0x7000
	EmmParaErrorInformation uint16 = 0x7001

	//table define

	EmmTableAuth          byte = 0x80
	EmmTableMessageSingle byte = 0x82
	EmmTableMessageArea   byte = 0x83
	EmmTableActive        byte = 0x82
	EmmTableMessage       byte = 0x83
)
