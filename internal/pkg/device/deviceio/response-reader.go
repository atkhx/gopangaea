package deviceio

import (
	"bytes"
	"io"
	"strings"
)

type responseReader struct {
	reader io.Reader
}

func NewResponseReader(reader io.Reader) *responseReader {
	return &responseReader{reader: reader}
}

func (d *responseReader) Read() ([]byte, error) {
	var result []byte

	for {
		readBuffer := make([]byte, 64)
		n, errRead := d.reader.Read(readBuffer)
		if errRead != nil {
			break
		}

		if n == 0 {
			break
		}

		result = append(result, readBuffer[:n]...)
	}
	return result, nil
}

func (d *responseReader) ReadWithSkipTails(command string, length int) ([]byte, error) {
	var checkTails = true
	var result []byte

	if idx := strings.Index(command, "\r"); idx > -1 {
		command = string([]byte(command)[:idx])
	}

	separator := append([]byte(command), byte(0x0d))

	skippedBytes := 0
	skippedBytesBuffer := []byte{}

	for checkTails || len(result) < length {
		readBuffer := make([]byte, 64)
		n, errRead := d.reader.Read(readBuffer)
		if errRead != nil {
			break
		}

		if n == 0 {
			break
		}

		readBuffer = readBuffer[:n]
		if checkTails {
			skippedBytesBuffer = append(skippedBytesBuffer, readBuffer...)
			skippedBytes = bytes.Index(skippedBytesBuffer, separator)
			if skippedBytes == -1 {
				continue
			}

			skippedBytes += len(separator)
			checkTails = false

			readBuffer = skippedBytesBuffer[skippedBytes:]
		}

		result = append(result, readBuffer...)
	}
	return result, nil
}
