// Package require is a wrapper over testify's require package.
package require

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Assertions adds assertions to testify's require lib.
type Assertions struct {
	*require.Assertions
	tt *testing.T
}

// New completely creates a new Assertions.
func New(tt *testing.T) *Assertions {
	return &Assertions{
		tt:         tt,
		Assertions: require.New(tt),
	}
}

// WantError is a that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   success := r.WantError(test.wantErr, err)
func (r *Assertions) WantError(wantErr bool, err error, msgAndArgs ...interface{}) {
	if h, ok := require.TestingT(r.tt).(tHelper); ok {
		h.Helper()
	}

	if wantErr {
		r.Error(err, msgAndArgs...)
	} else {
		r.NoError(err, msgAndArgs...)
	}
}

// tHelper is borrowed from a neat trick from testify.
type tHelper interface {
	Helper()
}
