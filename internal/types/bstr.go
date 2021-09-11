package types

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type Bstr struct {
	val *uint16
}

func (b *Bstr) Addr() uintptr {
	return uintptr(unsafe.Pointer(&b.val))
}

func (b *Bstr) String() string {
	return ole.BstrToString(b.val)
}

func (b *Bstr) Value() uintptr {
	return uintptr(unsafe.Pointer(b.val))
}
