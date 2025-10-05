package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		for {
			line := ""
			for {
				data := make([]byte, 8)
				n, err := f.Read(data)
				if err != nil {
					break
				}

				data = data[:n]
				if i := bytes.IndexByte(data, '\n'); i != -1 {
					line += string(data[:i])
					data = data[i+1:]
					out <- line
					line = ""
				}

				line += string(data)
			}

			if len(line) != 0 {
				out <- line
			}
		}
	}()

	return out
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
