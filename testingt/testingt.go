package testingt

// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
}

// Helper is borrowed from a neat trick from testify.
type Helper interface {
	Helper()
}
