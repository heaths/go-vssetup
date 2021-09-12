package formatting

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"golang.org/x/text/language"
)

func BenchmarkPrintString(b *testing.B) {
	w := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		printString(w, "a", a)
	}
}

func TestPrintString(t *testing.T) {
	w := &bytes.Buffer{}
	printString(w, "a", a)

	if w.String() != "a = 1\n" {
		t.Fatalf("got %s, expected a = 1", w.String())
	}
}

func BenchmarkPrintStringFunc(b *testing.B) {
	w := &bytes.Buffer{}
	p := newPrinter(w)
	for i := 0; i < b.N; i++ {
		p.printStringFunc(a)
	}
}

func TestPrintStringFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w)
	p.printStringFunc(a)

	if w.String() != "a = 1\n" {
		t.Fatalf(`got "%s", expected "a = 1"`, w.String())
	}
}

func TestPrintStringFunc_Method(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w)
	s := str{"1"}

	p.printStringFunc(s.String)

	if w.String() != "String = 1\n" {
		t.Fatalf(`got "%s", expected "String = 1"`, w.String())
	}
}

func TestPrintTimeFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w)
	p.printTimeFunc(b)

	if w.String() != "b = 2021-09-10 09:00:30 +0000 UTC\n" {
		t.Fatalf(`got "%s", expected "b = 2021-09-10 09:00:30 +0000 UTC"`, w.String())
	}
}

func TestPrintLocalizedStringFunc(t *testing.T) {
	w := &bytes.Buffer{}
	p := newPrinter(w)
	p.printLocalizedStringFunc(language.AmericanEnglish, c)

	if w.String() != "c = en-US\n" {
		t.Fatalf(`got "%s", expected "c = en-US"`, w.String())
	}
}

// Compared with printStringFunc, which offers suitable enough performance.
func printString(w io.Writer, name string, f func() (string, error)) {
	if s, err := f(); err == nil {
		fmt.Fprintf(w, "%s = %s\n", name, s)
	}
}

func a() (string, error) {
	return "1", nil
}

func b() (time.Time, error) {
	return time.Date(2021, 9, 10, 9, 00, 30, 0, time.UTC), nil
}

func c(locale language.Tag) (string, error) {
	return fmt.Sprint(locale), nil
}

type str struct {
	value string
}

func (s *str) String() (string, error) {
	return s.value, nil
}
