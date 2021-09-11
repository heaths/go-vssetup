package vssetup

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/interop"
	"github.com/heaths/go-vssetup/internal/types"
)

type query struct {
	v       *interop.ISetupConfiguration2
	err     error
	didInit bool
}

var q query

// Instances returns an array of Instance for Visual Studio 2017 and newer products.
// Set parameter all to true to enumerate all instances whether launchable or not.
func Instances(all bool) ([]Instance, error) {
	v, err := q.init()
	if err != nil {
		return nil, err
	} else if v == nil {
		// Assume no instances.
		return []Instance{}, nil
	}

	var e *interop.IEnumSetupInstances
	if all {
		e, err = v.EnumAllInstances()
		if err != nil {
			return nil, err
		}
	} else {
		e, err = v.EnumInstances()
		if err != nil {
			return nil, err
		}
	}

	instances := make([]Instance, 0)
	for {
		if insts, err := e.Next(1); err != nil {
			return nil, err
		} else if len(insts) == 1 {
			instance := Instance{insts[0]}
			instances = append(instances, instance)
		} else {
			break
		}
	}

	return instances, nil
}

// InstanceForCurrentProcess returns an *Instance for the current process or nil if none found.
func InstanceForCurrentProcess() (*Instance, error) {
	v, err := q.init()
	if err != nil {
		return nil, err
	} else if v == nil {
		return nil, nil
	}

	if inst, err := v.GetInstanceForCurrentProcess(); inst == nil || err != nil {
		return nil, err
	} else {
		return &Instance{inst}, nil
	}
}

// InstanceForPath returns an *Instance for the given path or nil if none found.
func InstanceForPath(path string) (*Instance, error) {
	v, err := q.init()
	if err != nil {
		return nil, err
	} else if v == nil {
		return nil, nil
	}

	bstr := types.NewBstr(path)
	if inst, err := v.GetInstanceForPath(bstr); inst == nil || err != nil {
		return nil, err
	} else {
		return &Instance{inst}, nil
	}
}

func (q *query) init() (*interop.ISetupConfiguration2, error) {
	if !q.didInit {
		if err := ole.CoInitialize(0); err != nil {
			if err.Error() == "" {
				err = ole.NewErrorWithSubError(ole.E_NOTIMPL, "not implemented", err)
			} else if e, ok := err.(*ole.OleError); ok && e.Code() == ole.E_NOTIMPL {
				// Likely not supported on the current platform, so don't try again.
				q.didInit = true
			}

			return nil, err
		}

		if unk, err := ole.CreateInstance(interop.CLSID_SetupConfiguration, interop.IID_ISetupConfiguration2); err != nil {
			if err.Error() == "" {
				q.err = ole.NewErrorWithSubError(ole.E_NOTIMPL, "not implemented", err)
			} else if e, ok := err.(*ole.OleError); ok && e.Code() == interop.REGDB_E_CLASSNOTREG {
				// No error. Assume no instances.
			} else {
				q.err = err
			}
		} else {
			q.v = (*interop.ISetupConfiguration2)(unsafe.Pointer(unk))
		}

		q.didInit = true
	}

	return q.v, q.err
}
