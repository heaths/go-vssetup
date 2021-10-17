package vssetup

import (
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/interop"
	"github.com/heaths/go-vssetup/internal/types"
	"golang.org/x/text/language"
)

type InstanceState uint32

const (
	None             InstanceState = 0
	Local            InstanceState = 1
	Registered       InstanceState = 2
	NoRebootRequired InstanceState = 4
	NoErrors         InstanceState = 8
	Complete         InstanceState = 4294967295
)

// Instance contains information about a Visual Studio 2017 or newer product.
type Instance struct {
	v  *interop.ISetupInstance
	v2 *interop.ISetupInstance2
}

func newInstance(v *interop.ISetupInstance) *Instance {
	i := &Instance{
		v: v,
	}

	runtime.SetFinalizer(i, (*Instance).Close)
	return i
}

// Close releases any resources used by this Instance immediately.
func (i *Instance) Close() error {
	if i.v != nil {
		// Call IUnknown.Release() but leave v assigned to avoid AV exceptions.
		i.v.Release()

		// Release ISetupInstance2 if initialized.
		if i.v2 != nil {
			i.v2.Release()
		}

		runtime.SetFinalizer(i, nil)
	}

	return nil
}

// InstanceID gets the unique, machine-specific ID for the Instance.
func (i *Instance) InstanceID() (string, error) {
	return getStringFunc(i.v.GetInstanceId)
}

// InstallDate gets the date the Instance was installed.
func (i *Instance) InstallDate() (time.Time, error) {
	return getTimeFunc(i.v.GetInstallDate)
}

// InstallationName gets the family name and version of the Instance.
func (i *Instance) InstallationName() (string, error) {
	return getStringFunc(i.v.GetInstallationName)
}

// InstallationPath gets the root path where the Instance was installed.
func (i *Instance) InstallationPath() (string, error) {
	return getStringFunc(i.v.GetInstallationPath)
}

// DisplayName gets the localized name of the Instance,
// or English if the name is not localized for the given locale.
func (i *Instance) DisplayName(locale language.Tag) (string, error) {
	return getLocalizedStringFunc(locale, i.v.GetDisplayName)
}

// Description gets the localized description of the Instance.
// or English if the name is not localized for the given locale.
func (i *Instance) Description(locale language.Tag) (string, error) {
	return getLocalizedStringFunc(locale, i.v.GetDescription)
}

// MakePath returns the combined Instance installation path with the given child path.
func (i *Instance) MakePath(path string) (string, error) {
	f := func() (*types.Bstr, error) {
		return i.v.ResolvePath(path)
	}
	return getStringFunc(f)
}

// State describes if the instance is complete or other combinations of InstanceState.
func (i *Instance) State() (InstanceState, error) {
	if err := i.initISetupInstance2(); err != nil {
		return None, err
	}

	if state, err := i.v2.GetState(); err != nil {
		return None, err
	} else {
		return InstanceState(state), nil
	}
}

// ProductPath gets the full path to the main executable, if defined.
func (i *Instance) ProductPath() (string, error) {
	if err := i.initISetupInstance2(); err != nil {
		return "", err
	}
	if s, err := getStringFunc(i.v2.GetProductPath); err != nil {
		return "", err
	} else {
		return i.MakePath(s)
	}
}

// IsLaunchable gets whether the instance can be launched.
func (i *Instance) IsLaunchable() (bool, error) {
	if err := i.initISetupInstance2(); err != nil {
		return false, err
	}
	return i.v2.IsLaunchable()
}

// IsComplete gets whether the instance has been completely installed.
func (i *Instance) IsComplete() (bool, error) {
	if err := i.initISetupInstance2(); err != nil {
		return false, err
	}
	return i.v2.IsComplete()
}

// IsRebootRequired gets whether the instance requires a reboot before launching.
func (i *Instance) IsRebootRequired() (bool, error) {
	state, err := i.State()
	return state&NoRebootRequired == 0, err
}

// EnginePath gets the path to the setup engine that installed this instance.
func (i *Instance) EnginePath() (string, error) {
	if err := i.initISetupInstance2(); err != nil {
		return "", err
	}
	return getStringFunc(i.v2.GetEnginePath)
}

func (i *Instance) initISetupInstance2() error {
	if i.v2 != nil {
		return nil
	}

	hr, _, _ := syscall.Syscall(
		i.v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(i.v)),
		uintptr(unsafe.Pointer(interop.IID_ISetupInstance2)),
		uintptr(unsafe.Pointer(&i.v2)),
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func getStringFunc(f func() (*types.Bstr, error)) (string, error) {
	if bstr, err := f(); err != nil {
		return "", err
	} else {
		defer bstr.Close()
		return bstr.String(), nil
	}
}

func getTimeFunc(f func() (*types.Filetime, error)) (time.Time, error) {
	if ft, err := f(); err != nil {
		return time.Time{}, err
	} else {
		return ft.Time(), nil
	}
}

func getLocalizedStringFunc(l language.Tag, f func(uint32) (*types.Bstr, error)) (string, error) {
	lcid := lcid(l)
	if bstr, err := f(lcid); err != nil {
		return "", err
	} else {
		defer bstr.Close()
		return bstr.String(), nil
	}
}
