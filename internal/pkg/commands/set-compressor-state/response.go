package set_compressor_state

// "00."
var successResponse = string([]byte{0x30, 0x30, 0x0a})

type Response struct {
	success bool
	content []byte
}

func ParseResponse(b []byte) (Response, error) {
	return Response{success: string(b) == successResponse}, nil
}

func (r Response) String() string {
	if r.success {
		return "success"
	}
	return "failed"
}
