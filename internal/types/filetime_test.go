package types

import (
	"testing"
	"time"
)

func Test_Time(t *testing.T) {
	ft := &Filetime{
		lowDateTime:  3577643008,
		highDateTime: 27111902,
	}

	want := time.Unix(0, 0)
	if got := ft.Time(); got != want {
		t.Errorf("Time() = %v, want = %v", got.UTC(), want.UTC())
	}
}
