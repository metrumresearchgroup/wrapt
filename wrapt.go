// Package wrapt wraps a *testing.T and adds functionality such as RunFatal,
// and Run with its own extended features.
package wrapt

import (
	"testing"

	"github.com/metrumresearchgroup/wrapt/assert"
	"github.com/metrumresearchgroup/wrapt/require"
)

// T is simply a wrap of *testing.T.
type T struct {
	// T is embedded so borrow all of the functionality for it, only
	// patching over in cases like Run, when necessary.
	*testing.T

	// A allows for assertions with a test error. The source of these
	// assertions is testify.
	A *assert.Assertions

	// A allows for requirements with a test failure. The source of
	// these assertions is also testify.
	R *require.Assertions

	// FatalHandler is a function that can be set to handle any failure
	// of a test where it is called, mainly in RunFatal.
	FatalHandler func(t *T, success bool, msgAndArgs ...interface{})
}

// WrapT takes a *testing.T and returns the equivalent *T from it. This is the
// entry point into all functionality, and should be done at the top of any
// *testing.T test to impart the new functionality.
func WrapT(tt *testing.T) *T {
	return &T{
		T: tt,
		A: assert.New(tt),
		R: require.New(tt),
		FatalHandler: func(t *T, success bool, msgAndArgs ...interface{}) {
			if !success {
				t.R.FailNow("fatal inner test failure", msgAndArgs...)
			}
		},
	}
}

// innerWrapT is an internal function that makes sure a new *T is returned
// from a *testing.T, but with the parent test's FatalHandler instead of
// a default one.
func (t *T) innerWrapT(tt *testing.T) *T {
	newT := WrapT(tt)
	newT.FatalHandler = t.FatalHandler

	return newT
}

// RunFatal is like .Run() but stops the outer test if the inner test fails.
// This is especially useful if a verification is critical to continuing.
func (t *T) RunFatal(name string, fn func(t *T)) {
	t.Helper()

	t.FatalHandler(t, t.Run(name, fn))
}

// Run implements the standard testing.T.Run() by wrapping *testing.T
// so the inner test has full access to our *T.
func (t *T) Run(name string, fn func(t *T)) (success bool) {
	t.Helper()

	return t.T.Run(name, t.wrapFn(fn))
}

// wrapFn wraps a function taking *T with a function that looks like it
// takes *testing.T so the testing.T.Run() can operate.
func (t *T) wrapFn(fn func(*T)) func(*testing.T) {
	return func(tt *testing.T) {
		fn(t.innerWrapT(tt))
	}
}
