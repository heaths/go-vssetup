//go:build !windows
// +build !windows

package interop

import (
	"github.com/heaths/go-vssetup/internal/errors"
	"github.com/heaths/go-vssetup/internal/types"
)

func (v *ISetupPackageReference) GetId() (*types.Bstr, error) { //nolint:stylecheck
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) GetVersion() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) GetChip() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) GetLanguage() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) GetBranch() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) GetType() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) GetUniqueId() (*types.Bstr, error) { //nolint:stylecheck
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) GetIsExtension() (bool, error) {
	return false, errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference) GetISetupPackageReference() (*ISetupPackageReference, error) {
	return nil, errors.NotImplemented(nil)
}
