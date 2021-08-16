// Package require is a wrapper over testify's require package.
package require

import (
	"github.com/stretchr/testify/require"
)

// Assertions adds assertions to testify's require lib.
type Assertions struct {
	*require.Assertions

	// We're using the require package's version so it keeps the
	// contract with our 3rd party.
	tt require.TestingT
}

// New completely creates a new Assertions.
func New(tt require.TestingT) *Assertions {
	return &Assertions{
		tt:         tt,
		Assertions: require.New(tt),
	}
}

// WantError is a that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   success := r.WantError(test.wantErr, err)
func (a *Assertions) WantError(wantErr bool, err error, msgAndArgs ...interface{}) {
	if h, ok := a.tt.(helper); ok {
		h.Helper()
	}

	if wantErr {
		a.Error(err, msgAndArgs...)
	} else {
		a.NoError(err, msgAndArgs...)
	}
}

// helper is borrowed from testify; it's copied here because
// it's called tHelper there, and is not exported.
type helper interface {
	Helper()
}
