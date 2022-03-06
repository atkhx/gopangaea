package get_mode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	tests := []struct {
		name          string
		data          []byte
		expected      Response
		expectedError bool
	}{
		{
			name: "valid response amtver.6.08.04.",
			data: []byte{0x61, 0x6d, 0x74, 0x76, 0x65, 0x72, 0x0d, 0x36, 0x2e, 0x30, 0x38, 0x2e, 0x30, 0x34, 0x0a},
			expected: Response{
				Major: 6,
				Minor: 8,
				Patch: 4,
			},
		},
		{
			name: "valid response amtver.0.00.00.",
			data: []byte{0x61, 0x6d, 0x74, 0x76, 0x65, 0x72, 0x0d, 0x30, 0x2e, 0x30, 0x30, 0x2e, 0x30, 0x30, 0x0a},
			expected: Response{
				Major: 0,
				Minor: 0,
				Patch: 0,
			},
		},
		{
			name: "valid response amtver.99.99.99.",
			data: []byte{0x61, 0x6d, 0x74, 0x76, 0x65, 0x72, 0x0d, 0x39, 0x39, 0x2e, 0x39, 0x39, 0x2e, 0x39, 0x39, 0x0a},
			expected: Response{
				Major: 99,
				Minor: 99,
				Patch: 99,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ParseResponse(tc.data)
			if (err != nil) != tc.expectedError {
				t.Errorf("ParseResponse() error = %v, expectedError %v", err, tc.expectedError)
				return
			}

			if !tc.expectedError {
				assert.Equal(t, tc.expected, actual)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
