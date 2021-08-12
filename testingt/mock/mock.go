// Package mock
package mock

type TestingT struct {
	Format        string
	Args          []interface{}
	ErrorfCalled  bool
	FailNowCalled bool
	HelperCalled  bool
}

func (m *TestingT) Errorf(format string, args ...interface{}) {
	m.Format = format
	m.Args = args
	m.ErrorfCalled = true
}

func (m *TestingT) FailNow() {
	m.FailNowCalled = true
}

func (m *TestingT) Helper() {
	m.HelperCalled = true
}
