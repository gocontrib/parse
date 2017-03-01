package parse

import (
	"bufio"
	"io"
	"strings"
	"sync"
)

func Lines(in io.Reader, parallel bool) chan string {
	output := make(chan string)

	go func() {
		scanner := bufio.NewScanner(in)
		var wg sync.WaitGroup

		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				continue
			}

			// some files could ended with zero bytes
			if line[0] == 0 {
				line = strings.Trim(line, string(rune(0)))
				if len(line) == 0 {
					continue
				}
			}

			if parallel {
				wg.Add(1)
				go func() {
					defer wg.Done()
					output <- line
				}()
			} else {
				output <- line
			}
		}

		if parallel {
			wg.Wait()
		}
		close(output)
	}()

	return output
}

func LinesChunked(in io.Reader, chunkSize int, parallel bool) chan []string {
	output := make(chan []string)

	go func() {
		chunk := make([]string, chunkSize)
		lines := Lines(in, parallel)
		i := 0

		for {
			line, ok := <-lines
			if !ok {
				break
			}

			chunk[i] = line
			i = i + 1

			if i >= chunkSize {
				i = 0
				output <- chunk[0:chunkSize]
			}
		}

		if i > 0 {
			output <- chunk[0:i]
		}

		close(output)
	}()

	return output
}
