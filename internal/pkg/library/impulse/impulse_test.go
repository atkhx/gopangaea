//go:build integration

package impulse

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test-data/origin.wav
var origin []byte

//go:embed test-data/origin.trimmed
var originTrimmed []byte

func TestImpulse_Source(t *testing.T) {
	impulse := New(origin)
	assert.Equal(t, origin, impulse.Source())
}

func TestImpulse_Trimmed(t *testing.T) {
	impulse := New(origin)
	trimmed, err := impulse.Trimmed()

	assert.NoError(t, err)
	assert.Equal(t, originTrimmed, trimmed)
}

func TestImpulse_IsValid(t *testing.T) {
	assert.NoError(t, New(origin).IsValid())
}
