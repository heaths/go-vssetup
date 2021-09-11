//go:build !windows
// +build !windows

package types

func NewBstr(s string) *Bstr {
	return &Bstr{}
}

func (b *Bstr) Close() error {
	return nil
}
