package cmd

import (
	"strconv"

	aphrodite "github.com/jonathon-chew/Aphrodite"
)

type Flags struct {
	Bytes        bool
	Lines        bool
	Speed        bool
	Elapsed      bool
	Rate         int
	SilentFinish bool
	HelpMenu     bool
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
		case "--Silent-finish", "-sf":
			options.SilentFinish = true
		case "--Rate", "-r":
			if len(commands[i:]) > 1 {
				number, ErrConvertingNumber := strconv.Atoi(commands[i+1])
				if ErrConvertingNumber != nil {
					i++
					continue
				}
				options.Rate = number
				i++
			}
		case "--help", "-h":
			// impliment help
			// os.Stdout.Write([]byte("Unfortunately help is not implimneted yet"))

			aphrodite.PrintInfo("--Bytes " + "-b " + "\n")
			aphrodite.PrintColour("Cyan", "Print bytes"+"\n")
			aphrodite.PrintByte([]byte("Prints out a running tally of bytes that have been proccesed\n\n"))

			aphrodite.PrintInfo("--Lines " + "-l " + "\n")
			aphrodite.PrintColour("Cyan", "Print lines"+"\n")
			aphrodite.PrintByte([]byte("Prints out a running tally of lines that have been proccesed\n\n"))

			aphrodite.PrintInfo("--Speed " + "-s " + "\n")
			aphrodite.PrintColour("Cyan", "Print speed"+"\n")
			aphrodite.PrintByte([]byte("Prints out a running tally of speed that the bytes are being passed through\n\n"))

			aphrodite.PrintInfo("--Elapsed " + "-e " + "--time " + "-t " + "\n")
			aphrodite.PrintColour("Cyan", "Print how much time has passed"+"\n")
			aphrodite.PrintByte([]byte("Prints out a running tally of how much time has passed since it started\n\n"))

			aphrodite.PrintInfo("--Rate " + "-r " + "\n")
			aphrodite.PrintColour("Cyan", "Print how fast data is being processed"+"\n")
			aphrodite.PrintByte([]byte("Determine how regularly to print the stats to the screen in milliseconds\n\n"))

			aphrodite.PrintInfo("If no options are passed in it will print Bytes, Lines, Speed and Time elpased and print to the screen every 200 milliseconds")

			options.HelpMenu = true
		}
	}

	return options
}
