// Package require is a wrapper over testify's require package. We can
// expand it with additional functionality with a minimum of code, so
// long as we don't collide with their definitions.
package require

import "github.com/stretchr/testify/require"

// Assertions adds assertions to testify's require lib.
type Assertions struct {
	// Embedding this to pass through to all functionality we don't add.
	*require.Assertions

	// We're using the require package's interface so it keeps the
	// contract with our 3rd party.
	tt require.TestingT
}

// New creates a new Assertions type wrapping testify with additional
// functions.
func New(tt require.TestingT) *Assertions {
	return &Assertions{
		// embedded fields carry the name of the underlying type.
		Assertions: require.New(tt),
		// we need tt to call helper.
		tt: tt,
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
