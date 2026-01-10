package main

import (
	"os"
	"log"
)

func main(){
	// fmt.Println("Hello World")
	bytesWritten, writeToStdErrError := os.Stderr.Write(os.Stdin)
	if writeToStdErrError != nil {
		log.Panic("ERROR: Writing to stderr")
	}

	log.Print("Bytes written: ", bytesWritten)
}
