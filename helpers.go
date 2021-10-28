package vssetup

import (
	"time"

	"github.com/heaths/go-vssetup/internal/types"
	"golang.org/x/text/language"
)

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
