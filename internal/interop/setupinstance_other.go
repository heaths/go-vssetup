//go:build !windows
// +build !windows

package interop

import (
	"github.com/heaths/go-vssetup/internal/errors"
	"github.com/heaths/go-vssetup/internal/types"
)

//nolint:stylecheck
func (v *ISetupInstance) GetInstanceId() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance) GetInstallDate() (*types.Filetime, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance) GetInstallationName() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance) GetInstallationPath() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance) GetInstallationVersion() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance) GetDisplayName(lcid uint32) (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance) GetDescription(lcid uint32) (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance) ResolvePath(relativePath string) (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance) ISetupInstance2(v2 **ISetupInstance2) error {
	return errors.NotImplemented(nil)
}

func (v *ISetupInstance2) GetState() (uint32, error) {
	return 0, errors.NotImplemented(nil)
}

func (v *ISetupInstance2) GetProduct() (*ISetupPackageReference, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance2) GetProductPath() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance2) IsLaunchable() (bool, error) {
	return false, errors.NotImplemented(nil)
}

func (v *ISetupInstance2) IsComplete() (bool, error) {
	return false, errors.NotImplemented(nil)
}

func (v *ISetupInstance2) GetProperties() (*ISetupPropertyStore, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupInstance2) GetEnginePath() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}
