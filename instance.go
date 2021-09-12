package vssetup

import (
	"runtime"
	"time"

	"github.com/heaths/go-vssetup/internal/interop"
	"github.com/heaths/go-vssetup/internal/types"
	"golang.org/x/text/language"
)

// Instance contains information about a Visual Studio 2017 or newer product.
type Instance struct {
	v *interop.ISetupInstance
}

func newInstance(v *interop.ISetupInstance) *Instance {
	i := &Instance{v}

	runtime.SetFinalizer(i, (*Instance).Close)
	return i
}

// Close releases any resources used by this Instance immediately.
func (i *Instance) Close() error {
	if i.v != nil {
		// Call IUnknown.Release() by leave v assigned to avoid AV exceptions.
		i.v.Release()
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
	if bstr, err := i.v.ResolvePath(path); err != nil {
		return "", err
	} else {
		return bstr.String(), nil
	}
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
