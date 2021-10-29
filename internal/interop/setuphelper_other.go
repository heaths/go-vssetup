//go:build !windows
// +build !windows

package interop

import (
	"github.com/heaths/go-vssetup/internal/errors"
)

func (v *ISetupHelper) ParseVersion(s string) (version uint64, err error) {
	return 0, errors.NotImplemented(nil)
}

func (v *ISetupHelper) ParseVersionRange(s string) (min, max uint64, err error) {
	return 0, 0, errors.NotImplemented(nil)
}
