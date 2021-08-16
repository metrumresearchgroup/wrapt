// Package assert is a wrapper over testify's assert package. We can
// expand it with additional functionality with a minimum of code, so
// long as we don't collide with their definitions.
package assert

import "github.com/stretchr/testify/assert"

// Assertions adds assertions to testify's assert lib.
type Assertions struct {
	// Embedding this to pass through to all functionality we don't add.
	*assert.Assertions

	// We're using the require package's interface so it keeps the
	// contract with our 3rd party.
	tt assert.TestingT
}

// New creates a new Assertions type wrapping testify with additional
// functions.
func New(tt assert.TestingT) *Assertions {
	return &Assertions{
		// embedded fields carry the name of the underlying type.
		Assertions: assert.New(tt),
		// we need tt to call helper.
		tt: tt,
	}
}

// WantError asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   success := a.WantError(test.wantErr, err)
func (a *Assertions) WantError(wantErr bool, err error, msgAndArgs ...interface{}) (success bool) {
	if h, ok := a.tt.(helper); ok {
		h.Helper()
	}

	if wantErr {
		return a.Error(err, msgAndArgs...)
	} else {
		return a.NoError(err, msgAndArgs...)
	}
}

// helper is borrowed from a neat trick from testify.
type helper interface {
	Helper()
}
