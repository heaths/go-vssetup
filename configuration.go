package vssetup

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/errors"
	"github.com/heaths/go-vssetup/internal/interop"
)

type query struct {
	v       *interop.ISetupConfiguration2
	err     error
	didInit bool
}

var q query

// Instances returns an array of Instance for Visual Studio 2017 and newer products.
// Set parameter all to true to enumerate all instances whether launchable or not.
func Instances(all bool) ([]*Instance, error) {
	v, err := q.init()
	if err != nil {
		return nil, err
	} else if v == nil {
		// Assume no instances.
		return nil, nil
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
	defer e.Release()

	instances := make([]*Instance, 0)
	for {
		if elems, err := e.Next(1); err != nil {
			return nil, err
		} else if len(elems) == 1 {
			instance := newInstance(elems[0])
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

	if instance, err := v.GetInstanceForCurrentProcess(); instance == nil || err != nil {
		if err, ok := err.(*errors.Error); ok && err.Code() == interop.E_NOTFOUND {
			return nil, nil
		}
		return nil, err
	} else {
		return newInstance(instance), nil
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

	if instance, err := v.GetInstanceForPath(path); instance == nil || err != nil {
		if err, ok := err.(*errors.Error); ok && err.Code() == interop.E_NOTFOUND {
			return nil, nil
		}
		return nil, err
	} else {
		return newInstance(instance), nil
	}
}

func (q *query) init() (*interop.ISetupConfiguration2, error) {
	// TODO: Consider runtime.SetFinalizer to Release() and CoUninitialize() and pass parent references to each child.
	if !q.didInit {
		if err := ole.CoInitialize(0); err != nil {
			if err.Error() == "" {
				err = errors.NotImplemented(err)
			} else if e, ok := err.(*errors.Error); ok && e.Code() == ole.E_NOTIMPL {
				// Likely not supported on the current platform, so don't try again.
				q.didInit = true
			}

			return nil, err
		}

		if unk, err := ole.CreateInstance(interop.CLSID_SetupConfiguration, interop.IID_ISetupConfiguration2); err != nil {
			if err.Error() == "" {
				q.err = errors.NotImplemented(err)
			} else if e, ok := err.(*errors.Error); ok && e.Code() == interop.REGDB_E_CLASSNOTREG {
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
