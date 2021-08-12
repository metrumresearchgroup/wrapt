package wrapt_test

import (
	"errors"
	"testing"

	"github.com/metrumresearchgroup/wrapt"
)

func Test_T_RunFatal(tt *testing.T) {
	t := wrapt.WrapT(tt)
	t.RunFatal("positive assertion without failure", func(t *wrapt.T) {
		t.A.True(true, "not true")
	})
}

func Test_T_ResultHandler(tt *testing.T) {
	t := wrapt.WrapT(tt)

	var success bool

	t.FatalHandler = func(t *wrapt.T, suc bool, msgAndArgs ...interface{}) {
		success = suc
	}

	t.RunFatal("antifail", func(t *wrapt.T) {
		// do nothing
	})

	t.A.True(success)
}

// Test_Assert_WantError is a wiring test, and does not walk any negative
// paths, because we're passing through a real testing.T.
func Test_Assert_WantError(tt *testing.T) {
	type Wanter interface {
		//
		WantError(wantErr bool, err error, msgAndArgs ...interface{}) (success bool)
	}
	t := wrapt.WrapT(tt)

	wanter := Wanter(t.A)

	t.A.True(wanter.WantError(false, nil))
	t.A.True(wanter.WantError(true, errors.New("error")))
}

// Test_Require_WantError is a wiring test, and does not walk any negative
// paths, because we're passing through a real testing.T.
func Test_Require_WantError(tt *testing.T) {
	type Wanter interface {
		WantError(wantErr bool, err error, msgAndArgs ...interface{})
	}

	t := wrapt.WrapT(tt)

	wanter := Wanter(t.R)

	wanter.WantError(false, nil)
	wanter.WantError(true, errors.New("error"))
}
