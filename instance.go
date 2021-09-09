package vssetup

import (
	"time"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/interop"
)

type Instance struct {
	v *interop.ISetupInstance
}

// Get the instance ID.
func (instance *Instance) InstanceId() (string, error) {
	if bstr, err := instance.v.GetInstanceId(); err != nil {
		return "", err
	} else {
		return ole.BstrToString(bstr), nil
	}
}

func (instance *Instance) InstallDate() (time.Time, error) {
	if ft, err := instance.v.GetInstallDate(); err != nil {
		return time.Time{}, err
	} else {
		return ft.Time(), nil
	}
}

func (instance *Instance) InstallationName() (string, error) {
	if bstr, err := instance.v.GetInstallationName(); err != nil {
		return "", err
	} else {
		return ole.BstrToString(bstr), nil
	}
}

func (instance *Instance) InstallationPath() (string, error) {
	if bstr, err := instance.v.GetInstallationPath(); err != nil {
		return "", err
	} else {
		return ole.BstrToString(bstr), nil
	}
}
