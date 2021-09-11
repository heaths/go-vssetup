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
			if got.Code() != ole.E_NOTIMPL {
				t.Errorf("NotImplemented() code = %#x, expected %#x", got.Code(), ole.E_NOTIMPL)
			} else if got.Error() != "Not implemented" {
				t.Errorf(`NotImplemented() = "%s", expected "%s`, got.Error(), "Not implemented")
			} else if (got.SubError() != nil) != tt.wantE {
				t.Errorf("NotImplemented() error = %v, expected? %t", got.SubError(), tt.wantE)
			}
		})
	}
}
