package assert_test

import (
	"errors"
	"testing"

	// vendorassert is the test library bound to our actual *testing.T.
	// Renaming the import distinguishes it from our own library
	// of the same name.
	vendorassert "github.com/stretchr/testify/assert"

	"github.com/metrumresearchgroup/wrapt/assert"
	"github.com/metrumresearchgroup/wrapt/assert/assertfakes"
)

// We're generating a mock for testify's TestingT in order to fail
// a test without actually failing the real test.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 github.com/stretchr/testify/assert.TestingT

func TestAssertions_WantError(t *testing.T) {
	type args struct {
		wantErr    bool
		err        error
		msgAndArgs []interface{}
	}
	// This represents the values passed on to Errorf by the underlying
	// framework.
	type expected struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		// the testmock function allows us to perform varying checks
		// on the result that aren't hard-coded in the test body,
		// expanding flexibility to tests.
		testMock func(r *vendorassert.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT)
		expected *expected
	}{
		{
			name: "success wantErr true",
			args: args{
				wantErr:    true,
				err:        errors.New("error"),
				msgAndArgs: []interface{}{"message", "hi"},
			},
			testMock: func(r *vendorassert.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT) {
				r.Equal(0, fakeTestingT.ErrorfCallCount())
			},
		},
		{
			name: "success wantErr false",
			args: args{
				wantErr:    false,
				err:        nil,
				msgAndArgs: []interface{}{"message", "hi"},
			},
			testMock: func(r *vendorassert.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT) {
				r.Equal(0, fakeTestingT.ErrorfCallCount())
			},
		},
		{
			name: "failure wantErr true",
			args: args{
				wantErr:    true,
				err:        nil,
				msgAndArgs: []interface{}{"message", "hi"},
			},
			testMock: func(r *vendorassert.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT) {
				r.Equal(1, fakeTestingT.ErrorfCallCount())
				format, args := fakeTestingT.ErrorfArgsForCall(0)
				r.Equal(exp.format, format)
				r.Equal(exp.args, args)
			},
			expected: &expected{
				format: "\n%s",
				args:   []interface{}{"\tError Trace:\t\n\tError:      \tAn error is expected but got nil.\n\tMessages:   \t[message hi]\n"},
			},
		},
		{
			name: "failure wantErr false",
			args: args{
				wantErr:    false,
				err:        errors.New("new"),
				msgAndArgs: []interface{}{"message", "hi"},
			},
			testMock: func(r *vendorassert.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT) {
				r.Equal(1, fakeTestingT.ErrorfCallCount())
				format, args := fakeTestingT.ErrorfArgsForCall(0)
				r.Equal(exp.format, format)
				r.Equal(exp.args, args)
			},
			expected: &expected{
				format: "\n%s",
				args:   []interface{}{"\tError Trace:\t\n\tError:      \tReceived unexpected error:\n\t            \tnew\n\tMessages:   \t[message hi]\n"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := vendorassert.New(t)
			// prevent a panic by ensuring the function is set before
			// using it.
			r.NotNil(test.testMock)

			// We're making a fake testing.T here conforming to testify's
			// model of it in the related package. Testify's assert and
			// require have different models for this.
			fakeTestingT := new(assertfakes.FakeTestingT)

			sut := assert.New(fakeTestingT)
			sut.WantError(test.args.wantErr, test.args.err, test.args.msgAndArgs)

			test.testMock(r, test.expected, fakeTestingT)
		})
	}
}
