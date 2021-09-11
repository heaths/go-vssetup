//go:build windows
// +build windows

package errors

import (
	"github.com/go-ole/go-ole"
)

func NotImplemented(err error) error {
	if err != nil {
		return ole.NewErrorWithSubError(ole.E_NOTIMPL, "", err)
	}

	return ole.NewError(ole.E_NOTIMPL)
}
