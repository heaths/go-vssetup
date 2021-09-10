package vssetup

import (
	"time"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/interop"
)

// Instance contains information about a Visual Studio 2017 or newer product.
type Instance struct {
	v *interop.ISetupInstance
}

// InstanceID gets the unique, machine-specific ID for the Instance.
func (instance *Instance) InstanceID() (string, error) {
	if bstr, err := instance.v.GetInstanceId(); err != nil {
		return "", err
	} else {
		return ole.BstrToString(bstr), nil
	}
}

// InstallDate gets the date the Instance was installed.
func (instance *Instance) InstallDate() (time.Time, error) {
	if ft, err := instance.v.GetInstallDate(); err != nil {
		return time.Time{}, err
	} else {
		return ft.Time(), nil
	}
}

// InstallationName gets the family name and version of the Instance.
func (instance *Instance) InstallationName() (string, error) {
	if bstr, err := instance.v.GetInstallationName(); err != nil {
		return "", err
	} else {
		return ole.BstrToString(bstr), nil
	}
}

// InstallationPath gets the root path where the Instance was installed.
func (instance *Instance) InstallationPath() (string, error) {
	if bstr, err := instance.v.GetInstallationPath(); err != nil {
		return "", err
	} else {
		return ole.BstrToString(bstr), nil
	}
}
