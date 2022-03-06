package deviceio

import "io"

type commandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *commandWriter {
	return &commandWriter{writer: writer}
}

func (w *commandWriter) Write(command string) error {
	_, err := w.writer.Write([]byte(command + "\r\n"))
	if err != nil {
		return err
	}
	return nil
}
