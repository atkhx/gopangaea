package set_poweramp_type

// "END."
var successResponse = string([]byte{0x45, 0x4e, 0x44, 0x0a})

type Response struct {
	success bool
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
