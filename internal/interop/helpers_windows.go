package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
)

func boolFunc(p, fn uintptr) (bool, error) {
	var b uint32
	if hr, _, _ := syscall.Syscall(fn, 2, p, uintptr(unsafe.Pointer(&b)), 0); hr != ole.S_OK {
		return false, ole.NewError(hr)
	}

	return b != 0, nil
}

func bstrFunc(p, fn uintptr) (*types.Bstr, error) {
	var bstr types.Bstr
	if hr, _, _ := syscall.Syscall(fn, 2, p, uintptr(bstr.Addr()), 0); hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return &bstr, nil
}

func localizedBstrFunc(p, fn uintptr, lcid uint32) (*types.Bstr, error) {
	var bstr types.Bstr
	if hr, _, _ := syscall.Syscall(fn, 2, p, uintptr(lcid), uintptr(bstr.Addr())); hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return &bstr, nil
}

func safeArrayFunc(p, fn uintptr) (array *ole.SafeArray, err error) {
	if hr, _, _ := syscall.Syscall(fn, 2, p, uintptr(unsafe.Pointer(&array)), 0); hr != ole.S_OK {
		err = ole.NewError(hr)
	}

	return
}

func stringArrayFunc(p, fn uintptr) ([]string, error) {
	var sa *ole.SafeArray
	if hr, _, _ := syscall.Syscall(fn, 2, p, uintptr(unsafe.Pointer(&sa)), 0); hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	array := ole.SafeArrayConversion{Array: sa}
	defer array.Release()

	return array.ToStringArray(), nil
}

func uint32Func(p, fn uintptr) (uint32, error) {
	var i uint32
	if hr, _, _ := syscall.Syscall(fn, 2, p, uintptr(unsafe.Pointer(&i)), 0); hr != ole.S_OK {
		return 0, ole.NewError(hr)
	}

	return i, nil
}
