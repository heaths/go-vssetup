package errors

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func Test_NotImplemented(t *testing.T) {
	tests := []struct {
		name  string
		err   error
		wantE bool
	}{
		{
			name:  "with error",
			err:   ole.NewError(ole.E_NOTIMPL),
			wantE: true,
		},
		{
			name: "without error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NotImplemented(tt.err)
			if err, ok := got.(*ole.OleError); !ok {
				t.Errorf("NotImplemented() expected *OleError, got %T", got)
			} else if err.Code() != ole.E_NOTIMPL {
				t.Errorf("NotImplemented() code = %#x, expected %#x", err.Code(), ole.E_NOTIMPL)
			} else if err.String() != "Not implemented" {
				t.Errorf(`NotImplemented() = "%s", expected "%s`, err.String(), "Not implemented")
			} else if (err.SubError() != nil) != tt.wantE {
				t.Errorf("NotImplemented() error = %v, expected? %t", err.SubError(), tt.wantE)
			}
		})
	}
}
