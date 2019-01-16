package main

import (
	"fmt"
	"log"
	"os"

	"github.com/takuyaohashi/go-wav"
)

func usage() {
	log.Fatal("Usage: waveparser [wave file name]")
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	w, err := wav.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	header := w.GetHeader()
	fmt.Printf("Sub Chunk Size \t= %d\n", header.SubChunkSize)
	fmt.Printf("Audio Format \t= %d\n", header.AudioFormat)
	fmt.Printf("Num Of Channels = %d\n", header.NumChannels)
	fmt.Printf("Sample Rate \t= %d\n", header.SampleRate)
	fmt.Printf("Byte Rate \t= %d\n", header.ByteRate)
	fmt.Printf("Block Align \t= %d\n", header.BlockAlign)
	fmt.Printf("Bit Per Sample \t= %d\n", header.BitsPerSample)
	fmt.Printf("Data Size \t= %d\n", header.SubChunk2Size)
}
