package set_lp_filter_state

type Response struct {
	success bool
	content []byte
}

func ParseResponse(b []byte) (Response, error) {
	if len(b) == 0 {
		return Response{success: true}, nil
	}
	return Response{success: false, content: b[:]}, nil
}

func (r Response) String() string {
	if r.success {
		return "success"
	}
	return "failed: " + string(r.content)
}
