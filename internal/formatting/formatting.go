package formatting

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/heaths/go-vssetup"
	"golang.org/x/text/language"
)

var (
	NameColor  = rgbColorFunc("9cdcfe")
	ValueColor = rgbColorFunc("ce9178")
)

type printer struct {
	nameFunc  func(string) string
	valueFunc func(string) string
	w         io.Writer
}

func PrintInstance(w io.Writer, i *vssetup.Instance, locale language.Tag) {
	p := newPrinter(w)

	p.printStringFunc(i.InstanceID)
	p.printTimeFunc(i.InstallDate)
	p.printStringFunc(i.InstallationName)
	p.printStringFunc(i.InstallationPath)
	p.printLocalizedStringFunc(locale, i.DisplayName)
	p.printLocalizedStringFunc(locale, i.Description)
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

func newPrinter(w io.Writer) *printer {
	if IsColorTerminal(w) {
		return &printer{
			nameFunc:  NameColor,
			valueFunc: ValueColor,
			w:         w,
		}
	} else {
		f := func(s string) string {
			return s
		}

		return &printer{
			nameFunc:  f,
			valueFunc: f,
			w:         w,
		}
	}
}

func (p *printer) printStringFunc(f func() (string, error)) {
	name := nameOf(f)
	if s, err := f(); err == nil {
		p.print(name, s)
	}
}

func (p *printer) printTimeFunc(f func() (time.Time, error)) {
	name := nameOf(f)
	if t, err := f(); err == nil {
		p.print(name, t.String())
	}
}

func (p *printer) printLocalizedStringFunc(l language.Tag, f func(language.Tag) (string, error)) {
	name := nameOf(f)
	if s, err := f(l); err == nil {
		p.print(name, s)
	}
}

func (p *printer) print(name, value string) {
	fmt.Fprintf(p.w, "%s = %s\n", p.nameFunc(name), p.valueFunc(value))
}

func rgbColorFunc(hex string) func(string) string {
	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	format := fmt.Sprintf("\033[38;2;%d;%d;%dm%%s\033[0m", r, g, b)

	return func(s string) string {
		return fmt.Sprintf(format, s)
	}
}
