package get_impulse_name

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	placeForImpulse := make([]byte, 71)
	copy(placeForImpulse, prefix)
	copy(placeForImpulse[71-len(suffix):], suffix)

	var validImpulseNameBytes []byte
	validImpulseName := "ma______________1960AV__________4x12_Vintage_30_Cond____Cap_0__.wav"
	validImpulseNameBytes = append(validImpulseNameBytes, prefix...)
	validImpulseNameBytes = append(validImpulseNameBytes, validImpulseName...)
	validImpulseNameBytes = append(validImpulseNameBytes, suffix...)

	var validShortNameBytes []byte
	validShortName := "some_cab.wav"
	validShortNameBytes = append(validShortNameBytes, prefix...)
	validShortNameBytes = append(validShortNameBytes, validShortName...)
	validShortNameBytes = append(validShortNameBytes, suffix...)

	var validEmptyBytes []byte
	validEmpty := "*"
	validEmptyBytes = append(validEmptyBytes, prefix...)
	validEmptyBytes = append(validEmptyBytes, validEmpty...)
	validEmptyBytes = append(validEmptyBytes, suffix...)

	tests := []struct {
		name          string
		data          []byte
		expected      Response
		expectedError bool
	}{
		{
			name:     "valid response",
			data:     validImpulseNameBytes,
			expected: Response{Name: validImpulseName},
		},
		{
			name:     "valid response",
			data:     validShortNameBytes,
			expected: Response{Name: validShortName},
		}, {
			name:     "valid response",
			data:     validEmptyBytes,
			expected: Response{Name: validEmpty},
		},
		{
			name:          "invalid response",
			data:          placeForImpulse,
			expectedError: true,
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
