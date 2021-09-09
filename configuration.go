package vssetup

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/interop"
)

var query = struct {
	v   *interop.ISetupConfiguration2
	err error
}{}

func init() {
	ole.CoInitialize(0)

	if unk, err := ole.CreateInstance(interop.CLSID_SetupConfiguration, interop.IID_ISetupConfiguration2); err != nil {
		if err.Error() == "" {
			query.err = ole.NewErrorWithSubError(ole.E_NOTIMPL, "not implemented", err)
		} else {
			query.err = err
		}
	} else {
		query.v = (*interop.ISetupConfiguration2)(unsafe.Pointer(unk))
	}
}

// Enumerates Visual Studio instances.
// Set parameter all to true to enumerate all instances whether launchable or not.
func Instances(all bool) ([]Instance, error) {
	if query.err != nil {
		return nil, query.err
	}

	var e *interop.IEnumSetupInstances
	var err error

	if all {
		e, err = query.v.EnumAllInstances()
		if err != nil {
			return nil, err
		}
	} else {
		e, err = query.v.EnumInstances()
		if err != nil {
			return nil, err
		}
	}

	instances := make([]Instance, 0)
	for {
		if _instances, err := e.Next(1); err != nil {
			return nil, err
		} else if len(_instances) == 1 {
			instance := Instance{_instances[0]}
			instances = append(instances, instance)
		} else {
			break
		}
	}

	return instances, nil
}
