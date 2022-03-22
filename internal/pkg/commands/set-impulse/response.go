package set_impulse

//ccEND..
var successResponse = string([]byte{0x63, 0x63, 0x45, 0x4e, 0x44, 0x0a, 0x0d})

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
