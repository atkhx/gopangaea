package deviceio

import (
	"bytes"
	"io"
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

	separator := append([]byte(command), byte(0x0d))

	skippedBytes := 0
	for len(result) < length {
		readBuffer := make([]byte, 64)
		n, errRead := d.reader.Read(readBuffer)
		if errRead != nil {
			break
		}

		if n == 0 {
			break
		}

		if checkTails {
			skippedBytes = bytes.Index(readBuffer[:n], separator)
			if skippedBytes == -1 {
				continue
			}

			checkTails = false
			readBuffer = readBuffer[skippedBytes:n]
			n -= skippedBytes
		}

		result = append(result, readBuffer[:n]...)
	}
	return result, nil
}
