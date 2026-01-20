package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	cli "github.com/jonathon-chew/go-pipepeek/internal/cli"
)

const (
	GigaByte = 1_000_000_000
	MegaByte = 1_000_000
	KiloByte = 1_000
)

func main() {

	var options cli.Flags

	if len(os.Args) > 1 {
		options = cli.Cmd(os.Args[1:])

		if options.HelpMenu {
			return
		}
	} else {
		options.Bytes = true
		options.Speed = true
		options.Elapsed = true
		options.SilentFinish = false
	}

	if options.Rate == 0 {
		options.Rate = 200
	} else {
		if !options.Bytes && !options.Speed && !options.Elapsed {
			options.Bytes = true
			options.Speed = true
			options.Elapsed = true
		}
	}

	// Constant variables
	buf := make([]byte, 32*1024)
	var totalBytes, totalLines, lastBytes int64
	var printBytes, printPerSecBytes string
	clearLine := []byte("\r\033[2K")

	// Grab information
	startTime := time.Now()

	// Variable - users maybe able to update this later
	renderEvery := time.Duration(options.Rate) * time.Millisecond

	// This means the stats will appear instantly not wait for the 200 m/s or what ever!
	lastRenderTime := startTime.Add(-renderEvery)
	now := time.Now()

	// Thie runs to notify users that something is happening
	os.Stderr.Write([]byte("waiting for input…"))

	for {
		n, err := os.Stdin.Read(buf)

		if n > 0 || err == io.EOF {
			totalBytes += int64(n)

			if options.Lines {
				for i := 0; i < n; i++ {
					if buf[i] == '\n' {
						totalLines++
					}
				}
			}

			// What time is it AFTER counting bites and new lines / any other logic
			now = time.Now()
			elapsed := now.Sub(startTime)

			// update status
			if now.Sub(lastRenderTime) > renderEvery || err == io.EOF {

				// Print everything to Stderr
				os.Stderr.Write(clearLine) // moves cursor back to the start of the line and clears the line
				// os.Stderr.Write([]byte("Read: "))

				if options.Bytes { // Processed bytes
					switch {
					case totalBytes >= GigaByte:
						printBytes = intToString(totalBytes/GigaByte) + "GBs"
					case totalBytes >= MegaByte:
						printBytes = intToString(totalBytes/MegaByte) + "MBs"
					case totalBytes >= KiloByte:
						printBytes = intToString(totalBytes/KiloByte) + "kBs"
					default:
						printBytes = intToString(totalBytes)
					}

					os.Stderr.Write([]byte(printBytes))
				}

				if options.Lines {

					// Human readable lines
					var buf [32]byte // enough for int64 with commas
					pos := len(buf)

					n := totalLines
					digitCount := 0

					if n == 0 {
						pos--
						buf[pos] = '0'
					} else {
						for n > 0 {
							if digitCount == 3 {
								pos--
								buf[pos] = ','
								digitCount = 0
							}

							pos--
							buf[pos] = byte('0' + n%10)
							n /= 10
							digitCount++
						}
					}

					lineLength := buf[pos:]

					// os.Stderr.Write([]byte(" | Lines: "))
					os.Stderr.Write([]byte(" | "))
					os.Stderr.Write(lineLength)
				}

				if options.Speed {
					//Calculations
					deltaBytes := totalBytes - lastBytes
					deltaTime := now.Sub(lastRenderTime)
					bytesPerSec := float64(deltaBytes) / deltaTime.Seconds()

					switch {
					case bytesPerSec >= GigaByte:
						printPerSecBytes = fmt.Sprintf("%.2f GB/s", bytesPerSec/(GigaByte))
					case bytesPerSec >= MegaByte:
						printPerSecBytes = fmt.Sprintf("%.2f MB/s", bytesPerSec/(MegaByte))
					case bytesPerSec >= KiloByte:
						printPerSecBytes = fmt.Sprintf("%.1f kB/s", bytesPerSec/KiloByte)
					default:
						printPerSecBytes = fmt.Sprintf("%.0f B/s", bytesPerSec)
					}

					// os.Stderr.Write([]byte(" | Rate: "))
					os.Stderr.Write([]byte(" | "))
					os.Stderr.Write([]byte(printPerSecBytes))
				}

				if options.Elapsed {
					// os.Stderr.Write([]byte(" | Elapsed: "))
					os.Stderr.Write([]byte(" | "))
					os.Stderr.Write([]byte(elapsed.Truncate(time.Second).String()))
				}
				// os.Stderr.Write([]byte("\n")) // left in for debugging, remove clear line and add this to check out puts

				// reset checked variables
				lastRenderTime = now
				lastBytes = totalBytes
			}

			// forward data to the pipe
			if _, werr := os.Stdout.Write(buf[:n]); werr != nil {
				log.Fatal(werr)
			}
		}

		if err == io.EOF {
			// Remove the stat line
			if options.SilentFinish {
				os.Stderr.Write(clearLine)
				os.Stderr.Write([]byte("\r"))
			} else {
				os.Stderr.Write([]byte("\n"))
			}
			break
		}

		if err != nil {
			log.Fatal(err)
		}
	}

}

func intToString(v int64) string {
	if v == 0 {
		return "0"
	}
	var b [20]byte
	i := len(b)
	for v > 0 {
		i--
		b[i] = byte('0' + v%10)
		v /= 10
	}
	return string(b[i:])
}
