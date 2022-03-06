package get_device

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
			name:     "valid response amtdev.04.END.",
			data:     []byte{0x61, 0x6d, 0x74, 0x64, 0x65, 0x76, 0x0d, 0x30, 0x34, 0x0a, 0x45, 0x4e, 0x44, 0x0a},
			expected: Response{Code: 4},
		},
		{
			name:     "valid response amtdev.00.END.",
			data:     []byte{0x61, 0x6d, 0x74, 0x64, 0x65, 0x76, 0x0d, 0x30, 0x30, 0x0a, 0x45, 0x4e, 0x44, 0x0a},
			expected: Response{Code: 0},
		},
		{
			name:     "valid response amtdev.99.END.",
			data:     []byte{0x61, 0x6d, 0x74, 0x64, 0x65, 0x76, 0x0d, 0x39, 0x39, 0x0a, 0x45, 0x4e, 0x44, 0x0a},
			expected: Response{Code: 99},
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
