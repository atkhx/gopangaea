package library

import (
	"github.com/atkhx/gopangaea/internal/pkg/library/impulse"
	"github.com/atkhx/gopangaea/internal/pkg/library/settings"
)

type Preset struct {
	Name     string
	Settings settings.Settings
	Impulse  *impulse.Impulse
}
