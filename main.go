package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Constant variables
	buf := make([]byte, 32*1024)
	var totalBytes, totalLines, lastBytes int64
	var printBytes, printPerSecBytes string
	clearLine := []byte("\r\033[2K")

	// Grab information
	lastRenderTime := time.Now()
	startTime := lastRenderTime

	// Variable - users maybe able to update this later
	renderEvery := 200 * time.Millisecond

	for {
		n, err := os.Stdin.Read(buf)

		if n > 0 {
			totalBytes += int64(n)

			for i := 0; i < n; i++ {
				if buf[i] == '\n' {
					totalLines++
				}
			}

			// What time is it AFTER counting bites and new lines / any other logic
			now := time.Now()
			elapsed := now.Sub(startTime)

			// update status
			if now.Sub(lastRenderTime) > renderEvery {
				//Calculations
				deltaBytes := totalBytes - lastBytes
				deltaTime := now.Sub(lastRenderTime)
				bytesPerSec := float64(deltaBytes) / deltaTime.Seconds()

				// Processed bytes
				switch {
				case totalBytes >= 1_000_000_000:
					printBytes = intToString(totalBytes/1_000_000_000) + "GBs"
				case totalBytes >= 1_000_000:
					printBytes = intToString(totalBytes/1_000_000) + "MBs"
				case totalBytes >= 1_000:
					printBytes = intToString(totalBytes/1_000) + "kBs"
				default:
					printBytes = intToString(totalBytes)
				}

				switch {
				case bytesPerSec >= 1_000_000_000:
					printPerSecBytes = fmt.Sprintf("%.2f GB/s", bytesPerSec/(1_000_000_000))
				case bytesPerSec >= 1_000_000:
					printPerSecBytes = fmt.Sprintf("%.2f MB/s", bytesPerSec/(1_000_000))
				case bytesPerSec >= 1_000:
					printPerSecBytes = fmt.Sprintf("%.1f kB/s", bytesPerSec/1_000)
				default:
					printPerSecBytes = fmt.Sprintf("%.0f B/s", bytesPerSec)
				}

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

				// Print everything to Stderr
				os.Stderr.Write(clearLine) // moves cursor back to the start of the line and clears the line
				os.Stderr.Write([]byte("Bytes read: "))
				os.Stderr.Write([]byte(printBytes))
				os.Stderr.Write([]byte(" | Lines read: "))
				os.Stderr.Write(lineLength)
				os.Stderr.Write([]byte(" | Rate: "))
				os.Stderr.Write([]byte(printPerSecBytes))
				os.Stderr.Write([]byte(" | Elapsed: "))
				os.Stderr.Write([]byte(elapsed.Truncate(time.Second).String()))
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
			os.Stderr.Write(clearLine)
			os.Stderr.Write([]byte("\r"))
			os.Stderr.Write([]byte("\n"))
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
