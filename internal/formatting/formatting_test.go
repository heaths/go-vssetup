package formatting

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"
)

func Benchmark_printString(b *testing.B) {
	w := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		printString(w, "a", a)
	}
}

func Test_printString(t *testing.T) {
	w := &bytes.Buffer{}
	printString(w, "a", a)

	if w.String() != "a = 1\n" {
		t.Fatalf("got %s, expected a = 1", w.String())
	}
}

func Benchmark_printStringFunc(b *testing.B) {
	w := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		printStringFunc(w, a)
	}
}

func Test_printStringFunc(t *testing.T) {
	w := &bytes.Buffer{}
	printStringFunc(w, a)

	if w.String() != "a = 1\n" {
		t.Fatalf(`got "%s", expected "a = 1"`, w.String())
	}
}

func Test_printStringFunc_Method(t *testing.T) {
	w := &bytes.Buffer{}
	s := str{"1"}
	printStringFunc(w, s.String)

	if w.String() != "String = 1\n" {
		t.Fatalf(`got "%s", expected "String = 1"`, w.String())
	}
}

func Test_printTimeFunc(t *testing.T) {
	w := &bytes.Buffer{}
	printTimeFunc(w, b)

	if w.String() != "b = 2021-09-10 09:00:30 +0000 UTC\n" {
		t.Fatalf(`got "%s", expected "b = 2021-09-10 09:00:30 +0000 UTC"`, w.String())
	}
}

func Test_printLocalizedStringFunc(t *testing.T) {
	w := &bytes.Buffer{}
	printLocalizedStringFunc(w, 1033, c)

	if w.String() != "c = 1033\n" {
		t.Fatalf(`got "%s", expected "c = 1033"`, w.String())
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

func c(lcid uint32) (string, error) {
	return fmt.Sprint(lcid), nil
}

type str struct {
	value string
}

func (s *str) String() (string, error) {
	return s.value, nil
}
