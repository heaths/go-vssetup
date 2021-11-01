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

func (v *ISetupFailedPackageReference) ISetupFailedPackageReference2(v2 **ISetupFailedPackageReference2) error {
	return errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference2) GetLogFilePath() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference2) GetDescription() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference2) GetSignature() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference2) GetDetails() ([]string, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference2) GetAffectedPackages() ([]*ISetupPackageReference, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference) ISetupFailedPackageReference3(v3 **ISetupFailedPackageReference3) error {
	return errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference3) GetAction() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupFailedPackageReference3) GetReturnCode() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) ISetupProductReference(pr **ISetupProductReference) error {
	return errors.NotImplemented(nil)
}

func (v *ISetupProductReference) GetIsInstalled() (bool, error) {
	return false, errors.NotImplemented(nil)
}

func (v *ISetupPackageReference) ISetupProductReference2(pr2 **ISetupProductReference2) error {
	return errors.NotImplemented(nil)
}

func (v *ISetupProductReference2) GetSupportsExtensions() (bool, error) {
	false, errors.NotImplemented(nil)
}
