package cmd

import (
	"os"
	"strconv"
)

type Flags struct {
	Bytes   bool
	Lines   bool
	Speed   bool
	Elapsed bool
	Rate    int
}

func Cmd(commands []string) Flags {

	var options Flags

	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "--Bytes", "-b":
			options.Bytes = true
		case "--Lines", "-l":
			options.Lines = true
		case "--Speed", "-s":
			options.Speed = true
		case "--Elapsed", "-e", "--time", "-t":
			options.Elapsed = true
		case "--Rate", "-r":
			if len(commands[i:]) > 1 {
				number, ErrConvertingNumber := strconv.Atoi(commands[i+1])
				if ErrConvertingNumber != nil {
					i++
					continue
				}
				options.Rate = number
			}
		case "--help", "-h":
			// impliment help
			os.Stdout.Write([]byte("Unfortunately help is not implimneted yet"))
		}
	}

	return options
}
