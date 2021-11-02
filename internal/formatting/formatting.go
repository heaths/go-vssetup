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
	"github.com/iancoleman/strcase"
	"golang.org/x/text/language"
)

var (
	// cSpell:ignore cdcfe
	NameColor  = rgbColorFunc("9cdcfe")
	ValueColor = func(i interface{}) string {
		return rgbColorFunc("ce9178")(fmt.Sprint(i))
	}
)

type printer struct {
	nameFunc  func(string) string
	valueFunc func(interface{}) string
	w         io.Writer
	opts      Options
}

type Includes int

const (
	Errors   Includes = 1
	Packages Includes = 2
)

type Options struct {
	Include Includes
	Locale  language.Tag
}

func PrintInstance(w io.Writer, i *vssetup.Instance, opts Options) {
	p := newPrinter(w, opts)

	p.printStringFunc("", i.InstanceID)
	p.printTimeFunc(i.InstallDate)
	p.printStringFunc("", i.InstallationName)
	p.printStringFunc("", i.InstallationPath)
	p.printStringFunc("", i.ProductPath)
	p.printStateFunc(i.State)
	p.printBoolFunc("", i.IsLaunchable)
	p.printBoolFunc("", i.IsComplete)
	p.printBoolFunc("", i.IsPrerelease)
	p.printBoolFunc("", i.IsRebootRequired)
	p.printLocalizedStringFunc(opts.Locale, i.DisplayName)
	p.printLocalizedStringFunc(opts.Locale, i.Description)
	p.printStringFunc("", i.EnginePath)
	p.printMapFunc("catalog_", i.CatalogInfo)
	p.printMapFunc("properties_", i.Properties)

	if opts.Include&Packages != 0 {
		if product, err := i.Product(); err == nil {
			defer product.Close()
			p.printProductReference("product_", product)
		}

		if packages, err := i.Packages(); err == nil {
			for idx, pkg := range packages {
				defer pkg.Close()
				prefix := fmt.Sprintf("package_%04d_", idx)
				p.printPackageReference(prefix, pkg)
			}
		}
	}

	if opts.Include&Errors != 0 {
		if errorState, err := i.ErrorState(); errorState != nil && err == nil {
			defer errorState.Close()
			p.printErrorState("errors_", errorState)
		}
	}
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

func newPrinter(w io.Writer, opts Options) *printer {
	if IsColorTerminal(w) {
		return &printer{
			nameFunc:  NameColor,
			valueFunc: ValueColor,
			w:         w,
			opts:      opts,
		}
	} else {
		return &printer{
			nameFunc: func(s string) string {
				return s
			},
			valueFunc: func(i interface{}) string {
				return fmt.Sprint(i)
			},
			w:    w,
			opts: opts,
		}
	}
}

func (p *printer) printBoolFunc(prefix string, f func() (bool, error)) {
	name := nameOf(f)
	if b, err := f(); err == nil {
		p.print(prefix, name, or(b, "1", "0"))
	}
}

func (p *printer) printStateFunc(f func() (vssetup.InstanceState, error)) {
	name := nameOf(f)
	if state, err := f(); err == nil {
		p.print("", name, strconv.FormatUint(uint64(state), 10))
	}
}

func (p *printer) printStringFunc(prefix string, f func() (string, error)) {
	name := nameOf(f)
	if s, err := f(); s != "" && err == nil {
		p.print(prefix, name, s)
	}
}

func (p *printer) printStringArrayFunc(prefix string, f func() ([]string, error)) {
	name := nameOf(f)
	if arr, err := f(); err == nil && len(arr) > 0 {
		p.print(prefix, name, strings.Join(arr, " "))
	}
}

func (p *printer) printTimeFunc(f func() (time.Time, error)) {
	name := nameOf(f)
	if t, err := f(); err == nil {
		p.print("", name, t.String())
	}
}

func (p *printer) printLocalizedStringFunc(l language.Tag, f func(language.Tag) (string, error)) {
	name := nameOf(f)
	if s, err := f(l); s != "" && err == nil {
		p.print("", name, s)
	}
}

func (p *printer) printErrorState(prefix string, errorState *vssetup.ErrorState) {
	if p.opts.Include&Packages != 0 {
		if packages, err := errorState.FailedPackages(); err == nil {
			for i, ref := range packages {
				defer ref.Close()
				// cSpell:ignore sfailed
				pre := fmt.Sprintf("%sfailed_%04d_", prefix, i)
				p.printFailedPackageReference(pre, ref)
			}
		}
		if packages, err := errorState.SkippedPackages(); err == nil {
			for i, ref := range packages {
				defer ref.Close()
				// cSpell:ignore sskipped
				pre := fmt.Sprintf("%sskipped_%04d_", prefix, i)
				p.printPackageReference(pre, ref)
			}
		}
	}
	p.printStringFunc(prefix, errorState.ErrorLogPath)
	p.printStringFunc(prefix, errorState.LogPath)
}

func (p *printer) printPackageReference(prefix string, ref *vssetup.PackageReference) {
	p.printStringFunc(prefix, ref.ID)
	p.printStringFunc(prefix, ref.Version)
	p.printStringFunc(prefix, ref.Chip)
	p.printStringFunc(prefix, ref.Language)
	p.printStringFunc(prefix, ref.Branch)
	p.printStringFunc(prefix, ref.Type)
	p.printStringFunc(prefix, ref.UniqueID)

	if ok, _ := ref.IsExtension(); ok {
		p.printBoolFunc(prefix, ref.IsExtension)
	}
}

func (p *printer) printFailedPackageReference(prefix string, ref *vssetup.FailedPackageReference) {
	p.printPackageReference(prefix, &ref.PackageReference)
	p.printStringFunc(prefix, ref.LogFilePath)
	p.printStringFunc(prefix, ref.Description)
	p.printStringFunc(prefix, ref.Action)
	p.printStringFunc(prefix, ref.ReturnCode)
	p.printStringArrayFunc(prefix, ref.Details)
}

func (p *printer) printProductReference(prefix string, ref *vssetup.ProductReference) {
	p.printPackageReference(prefix, &ref.PackageReference)
	p.printBoolFunc(prefix, ref.IsInstalled)
	p.printBoolFunc(prefix, ref.SupportsExtensions)
}

func (p *printer) printMapFunc(prefix string, f func() (map[string]interface{}, error)) {
	if m, err := f(); err == nil {
		names := make([]string, 0, len(m))
		for name := range m {
			names = append(names, name)
		}

		sort.Strings(names)

		for _, name := range names {
			p.print(prefix, name, m[name])
		}
	}
}

func (p *printer) print(prefix, name string, value interface{}) {
	name = strcase.ToLowerCamel(name)
	name = p.nameFunc(prefix + name)

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
