package formatting

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/heaths/go-vssetup"
	"golang.org/x/text/language"
)

func BenchmarkPrintString(b *testing.B) {
	w := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		printString(w, "a", stringFunc)
	}
}

func TestPrintString(t *testing.T) {
	w := &bytes.Buffer{}
	printString(w, "a", stringFunc)

	want := "a = 1\n"
	if w.String() != want {
		t.Fatalf("got %s, expected %q", w.String(), want)
	}
}

func BenchmarkPrintStringFunc(b *testing.B) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	for i := 0; i < b.N; i++ {
		p.printStringFunc("", stringFunc)
	}
}

func TestPrintStringFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.printStringFunc("", stringFunc)

	want := "stringFunc = 1\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrintStringFunc_Empty(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.printStringFunc("", stringFuncEmpty)

	want := ""
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrintStringFunc_Method(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	s := str{"1"}

	p.printStringFunc("", s.String)

	want := "string = 1\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrintStringArrayFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.printStringArrayFunc("", stringArrayFunc)

	want := "stringArrayFunc = a b\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrintBoolFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.printBoolFunc("", boolFunc(false))
	p.printBoolFunc("", boolFunc(true))

	want := "func1 = 0\nfunc2 = 1\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrintStateFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.printStateFunc(stateFunc)

	want := "stateFunc = 4294967295\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrintTimeFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.printTimeFunc(timeFunc)

	want := "timeFunc = 2021-09-10 09:00:30 +0000 UTC\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrintLocalizedStringFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.printLocalizedStringFunc(language.AmericanEnglish, localizedFunc)

	want := "localizedFunc = en-US\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrintMapFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.printMapFunc("prefix_", mapFunc)

	want := "prefix_a = 1\nprefix_b = 2\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

func TestPrint_CamelCase(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w, Options{})
	p.print("", "CamelCase", 1)
	p.print("prefix_", "CamelCase", 2)

	want := "camelCase = 1\nprefix_camelCase = 2\n"
	if w.String() != want {
		t.Fatalf(`got %q, expected %q`, w.String(), want)
	}
}

// Compared with printStringFunc, which offers suitable enough performance.
func printString(w io.Writer, name string, f func() (string, error)) {
	if s, err := f(); err == nil {
		fmt.Fprintf(w, "%s = %s\n", name, s)
	}
}

func stringFunc() (string, error) {
	return "1", nil
}

func stringFuncEmpty() (string, error) {
	return "", nil
}

func stringArrayFunc() ([]string, error) {
	return []string{"a", "b"}, nil
}

func boolFunc(b bool) func() (bool, error) {
	return func() (bool, error) {
		return b, nil
	}
}

func stateFunc() (vssetup.InstanceState, error) {
	return vssetup.Complete, nil
}

func timeFunc() (time.Time, error) {
	return time.Date(2021, 9, 10, 9, 00, 30, 0, time.UTC), nil
}

func localizedFunc(locale language.Tag) (string, error) {
	return fmt.Sprint(locale), nil
}

func mapFunc() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"a": 1,
		"b": "2",
	}
	return m, nil
}

type str struct {
	value string
}

func (s *str) String() (string, error) {
	return s.value, nil
}
