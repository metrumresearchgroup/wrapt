// Package assert is a wrapper over testify's assert package.
package assert

import (
	"github.com/stretchr/testify/assert"
)

// Assertions adds assertions to testify's assert lib.
type Assertions struct {
	*assert.Assertions
	tt assert.TestingT
}

// New properly creates an Assertions.
func New(tt assert.TestingT) *Assertions {
	return &Assertions{
		tt:         tt,
		Assertions: assert.New(tt),
	}
}

// WantError asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   success := a.WantError(test.wantErr, err)
func (a *Assertions) WantError(wantErr bool, err error, msgAndArgs ...interface{}) (success bool) {
	if h, ok := a.tt.(tHelper); ok {
		h.Helper()
	}

	if wantErr {
		return a.Error(err, msgAndArgs...)
	} else {
		return a.NoError(err, msgAndArgs...)
	}
}

// tHelper is borrowed from a neat trick from testify.
type tHelper interface {
	Helper()
}
