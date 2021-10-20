package formatting

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/heaths/go-vssetup"
	"golang.org/x/text/language"
)

var (
	NameColor  = rgbColorFunc("9cdcfe")
	ValueColor = func(i interface{}) string {
		return rgbColorFunc("ce9178")(fmt.Sprint(i))
	}
)

type printer struct {
	nameFunc  func(string) string
	valueFunc func(interface{}) string
	w         io.Writer
}

func PrintInstance(w io.Writer, i *vssetup.Instance, locale language.Tag) {
	p := newPrinter(w)

	p.printStringFunc(i.InstanceID)
	p.printTimeFunc(i.InstallDate)
	p.printStringFunc(i.InstallationName)
	p.printStringFunc(i.InstallationPath)
	p.printStringFunc(i.ProductPath)
	p.printStateFunc(i.State)
	p.printBoolFunc(i.IsLaunchable)
	p.printBoolFunc(i.IsComplete)
	p.printBoolFunc(i.IsRebootRequired)
	p.printLocalizedStringFunc(locale, i.DisplayName)
	p.printLocalizedStringFunc(locale, i.Description)
	p.printStringFunc(i.EnginePath)
	p.printMapFunc("properties_", i.Properties)
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
		return &printer{
			nameFunc: func(s string) string {
				return s
			},
			valueFunc: func(i interface{}) string {
				return fmt.Sprint(i)
			},
			w: w,
		}
	}
}

func (p *printer) printBoolFunc(f func() (bool, error)) {
	name := nameOf(f)
	if b, err := f(); err == nil {
		p.print(name, or(b, "1", "0"))
	}
}

func (p *printer) printStateFunc(f func() (vssetup.InstanceState, error)) {
	name := nameOf(f)
	if state, err := f(); err == nil {
		p.print(name, strconv.FormatUint(uint64(state), 10))
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

func (p *printer) printMapFunc(prefix string, f func() (map[string]interface{}, error)) {
	if m, err := f(); err == nil {
		names := make([]string, 0, len(m))
		for name := range m {
			names = append(names, name)
		}

		sort.Strings(names)

		for _, name := range names {
			p.print(prefix+name, m[name])
		}
	}
}

func (p *printer) print(name string, value interface{}) {
	fmt.Fprintf(p.w, "%s = %s\n", p.nameFunc(name), p.valueFunc(value))
}

func or(b bool, x, y string) string {
	if b {
		return x
	}
	return y
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
