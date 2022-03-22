package deviceio

// "END."
var responseWithEND = []byte{0x45, 0x4e, 0x44, 0x0a}

// "00."
var responseWith00 = []byte{0x30, 0x30, 0x0a}

func NewResponseBoolWithCustomEnd(suffix []byte) *ResponseBool {
	return &ResponseBool{
		expected:  suffix,
		maxLength: len(suffix),
	}
}
func NewResponseBoolWithoutEnd() *ResponseBool {
	return &ResponseBool{
		expected:  []byte{},
		maxLength: 0,
	}
}

func NewResponseBoolWithEnd() *ResponseBool {
	return &ResponseBool{
		expected:  responseWithEND,
		maxLength: len(responseWithEND),
	}
}

func NewResponseBoolWithZeros() *ResponseBool {
	return &ResponseBool{
		expected:  responseWith00,
		maxLength: len(responseWith00),
	}
}

type ResponseBool struct {
	maxLength int
	expected  []byte
	success   bool
}

func (r *ResponseBool) GetLength() int {
	return r.maxLength
}

func (r *ResponseBool) Parse(actual []byte) error {
	r.success = string(actual) == string(r.expected)
	return nil
}

func (r ResponseBool) Success() bool {
	return r.success
}

func (r ResponseBool) String() string {
	if r.success {
		return "success"
	}
	return "failed"
}
