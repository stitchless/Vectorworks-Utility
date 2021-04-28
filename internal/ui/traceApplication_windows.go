package ui

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

func runApplication(ch chan []byte, targetFile string) {
	cmd := exec.Command(targetFile)

	outReader, outWriter := io.Pipe()
	errReader, errWriter := io.Pipe()

	cmd.Stdout = outWriter
	cmd.Stderr = errWriter

	//mw := io.MultiWriter(outWriter, errWriter, &buffer)
	mr := io.MultiReader(outReader, errReader, &buffer)

	// Is the transfer the output of the application, through the channel.

	go func() {
		reader := bufio.NewReader(mr)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			ch <- line
			ch <- []byte{'\n'}
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Panicln(err)
	}
}
