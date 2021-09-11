package formatting

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/heaths/go-vssetup"
)

func PrintInstance(w io.Writer, i *vssetup.Instance) {
	lcid := vssetup.Locale()

	printStringFunc(w, i.InstanceID)
	printTimeFunc(w, i.InstallDate)
	printStringFunc(w, i.InstallationName)
	printStringFunc(w, i.InstallationPath)
	printLocalizedStringFunc(w, lcid, i.DisplayName)
	printLocalizedStringFunc(w, lcid, i.Description)
}

func nameOf(f interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	if idx := strings.LastIndex(name, "."); idx >= 0 {
		name = name[idx+1:]
	}

	// Strip decorators from method names.
	if idx := strings.Index(name, "-"); idx >= 0 {
		name = name[:idx]
	}

	return name
}

func printStringFunc(w io.Writer, f func() (string, error)) {
	name := nameOf(f)
	if s, err := f(); err == nil {
		fmt.Fprintf(w, "%s = %s\n", name, s)
	}
}

func printTimeFunc(w io.Writer, f func() (time.Time, error)) {
	name := nameOf(f)
	if t, err := f(); err == nil {
		fmt.Fprintf(w, "%s = %s\n", name, t)
	}
}

func printLocalizedStringFunc(w io.Writer, lcid uint32, f func(uint32) (string, error)) {
	name := nameOf(f)
	if t, err := f(lcid); err == nil {
		fmt.Fprintf(w, "%s = %s\n", name, t)
	}
}
