package set_equalizer_mixer

var successResponse = string([]byte{})

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
