package deviceio

import "io"

type commandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *commandWriter {
	return &commandWriter{writer: writer}
}

func (w *commandWriter) Write(command string) error {
	//buffer := []byte(command + "\r\n")
	//for i := 0; i < len(buffer); i += 501 {
	//	n := 501
	//	if i+n > len(buffer) {
	//		n = len(buffer) - i
	//	}
	//	_, err := w.writer.Write(buffer[i : i+n])
	//	if err != nil {
	//		return err
	//	}
	//}
	_, err := w.writer.Write([]byte(command + "\r\n"))
	if err != nil {
		return err
	}
	return nil
}
