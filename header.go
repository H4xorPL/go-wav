package wav

import (
	"fmt"
)

const (
	headerSize = 44
)

type AudioFormat int

const (
	AudioFormatPCM AudioFormat = iota
	AudioFormatBitstream
	AudioFormatInvalid
)
const (
	PCMAudioFormat = "PCM"
	BitStreamAudioFormat = "BitStream"
)

func (af AudioFormat) String() string {
	switch af {
	case AudioFormatPCM:
		return PCMAudioFormat
	case AudioFormatBitstream:
		return BitStreamAudioFormat
	default:
		return "Error"
	}
}

func GetAudioFormat(format string) (AudioFormat,error){
	switch format {
	case PCMAudioFormat:
		return AudioFormatPCM,nil
	case BitStreamAudioFormat:
		return AudioFormatBitstream,nil
	}
	return AudioFormatInvalid, fmt.Errorf("could not get audio format for: %s", format)
}

// WaveHeader is wave header struct
type waveHeader struct {
	chunkSize     uint32
	subChunkSize  uint32
	audioFormat   uint16
	numChannels   uint16
	sampleRate    uint32
	byteRate      uint32
	blockAlign    uint16
	bitsPerSample uint16
	subChunk2Size uint32
}
