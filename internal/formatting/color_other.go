//go:build !windows
// +build !windows

package formatting

import (
	"io"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

func IsColorTerminal(w io.Writer) bool {
	if f, ok := w.(*os.File); ok {
		if isatty.IsTerminal(f.Fd()) {
			term := os.Getenv("TERM")
			colorterm := os.Getenv("COLORTERM")

			return strings.Contains(term, "256") ||
				strings.Contains(term, "truecolor") ||
				strings.Contains(colorterm, "256") ||
				strings.Contains(colorterm, "truecolor")
		}
	}

	return false
}
