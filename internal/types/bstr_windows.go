//go:build windows
// +build windows

package types

import (
	"runtime"

	"github.com/heaths/go-vssetup/internal/windows"
)

func NewBstr(s string) *Bstr {
	bstr := &Bstr{
		val: windows.SysAllocString(s),
	}

	runtime.SetFinalizer(bstr, (*Bstr).Close)
	return bstr
}

func (b *Bstr) Close() error {
	if b.val != nil {
		if err := windows.SysFreeString(b.val); err != nil {
			return err
		}

		b.val = nil
		runtime.SetFinalizer(b, nil)
	}

	return nil
}
