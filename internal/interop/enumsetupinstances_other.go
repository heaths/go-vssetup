//go:build !windows
// +build !windows

package interop

import "github.com/heaths/go-vssetup/internal/errors"

func (v *IEnumSetupInstances) Next(celt uint32) ([]*ISetupInstance, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *IEnumSetupInstances) Skip(celt uint32) error {
	return errors.NotImplemented(nil)
}

func (v *IEnumSetupInstances) Reset() error {
	return errors.NotImplemented(nil)
}

func (v *IEnumSetupInstances) Clone() (*IEnumSetupInstances, error) {
	return nil, errors.NotImplemented(nil)
}
