package get_bank

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
			name:     "valid response with zeros",
			data:     []byte{0x67, 0x62, 0x0d, 0x30, 0x33, 0x30, 0x30, 0x0a},
			expected: Response{Bank: 3, Preset: 0},
		},
		{
			name:          "invalid response by 2 byte",
			data:          []byte{0x67, 0x62, 0x2e, 0x30, 0x33, 0x30, 0x30, 0x0a},
			expectedError: true,
		},
		{
			name:          "invalid response by 7 byte",
			data:          []byte{0x67, 0x62, 0x0d, 0x30, 0x33, 0x30, 0x30, 0x2e},
			expectedError: true,
		},
		{
			name:     "valid response with 0, 0",
			data:     []byte{0x67, 0x62, 0x0d, 0x30, 0x30, 0x30, 0x30, 0x0a},
			expected: Response{Bank: 0, Preset: 0},
		},
		{
			name:     "valid response with 9, 9",
			data:     []byte{0x67, 0x62, 0x0d, 0x30, 0x39, 0x30, 0x39, 0x0a},
			expected: Response{Bank: 9, Preset: 9},
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
