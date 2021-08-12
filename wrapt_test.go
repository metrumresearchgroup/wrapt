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

func Test_Assert_WantError(tt *testing.T) {
	t := wrapt.WrapT(tt)

	t.A.WantError(false, nil)
	t.A.WantError(true, errors.New("error"))
}

func Test_Require_WantError(tt *testing.T) {
	t := wrapt.WrapT(tt)

	t.R.WantError(false, nil)
	t.R.WantError(true, errors.New("error"))
}
