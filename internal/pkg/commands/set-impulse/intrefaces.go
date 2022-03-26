//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock_$GOFILE
package set_impulse

type Impulse interface {
	IsValid() error
	Trimmed() ([]byte, error)
	Source() []byte
}
