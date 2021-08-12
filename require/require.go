// Package require is a wrapper over testify's require package.
package require

import (
	"github.com/stretchr/testify/require"

	"github.com/metrumresearchgroup/wrapt/testingt"
)

// Assertions adds assertions to testify's require lib.
type Assertions struct {
	*require.Assertions
	tt testingt.TestingT
}

// New completely creates a new Assertions.
func New(tt testingt.TestingT) *Assertions {
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
	if h, ok := r.tt.(testingt.Helper); ok {
		h.Helper()
	}

	if wantErr {
		r.Error(err, msgAndArgs...)
	} else {
		r.NoError(err, msgAndArgs...)
	}
}
