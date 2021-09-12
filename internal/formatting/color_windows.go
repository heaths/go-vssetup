//go:build windows
// +build windows

package formatting

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
	"golang.org/x/sys/windows"
)

func IsColorTerminal(w io.Writer) bool {
	if f, ok := w.(*os.File); ok {
		if isatty.IsTerminal(f.Fd()) {
			return enableVirtualTerminalProcessing(f)
		}
	}

	return false
}

func enableVirtualTerminalProcessing(f *os.File) bool {
	stdout := windows.Handle(f.Fd())

	var err error
	var originalMode uint32

	if err = windows.GetConsoleMode(stdout, &originalMode); err == nil {
		err = windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	}

	return err == nil
}
