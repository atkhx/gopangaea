package set_impulse

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCommand_GetCommandTemporary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trimmed := []byte(`some trimmed`)

	impulse := NewMockImpulse(ctrl)
	impulse.EXPECT().Trimmed().Return(trimmed, nil)

	command := New("some name", false, impulse)
	actual := command.GetCommand()

	assert.Equal(t, "cc s 1\r736f6d65207472696d6d6564\r", actual)
}

func TestCommand_GetCommandPersist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	origin := []byte(`some origin`)

	extractor := NewMockImpulse(ctrl)
	extractor.EXPECT().Source().Return(origin)

	command := New("some name", true, extractor)
	actual := command.GetCommand()

	assert.Equal(t, "cc some name 0\r736f6d65206f726967696e\r", actual)
}
