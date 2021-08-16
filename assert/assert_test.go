package assert_test

import (
	"errors"
	"testing"

	vendorrequire "github.com/stretchr/testify/require"

	"github.com/metrumresearchgroup/wrapt/assert"
	"github.com/metrumresearchgroup/wrapt/assert/assertfakes"
)

func TestAssertions_WantError(t *testing.T) {
	type args struct {
		wantErr    bool
		err        error
		msgAndArgs []interface{}
	}
	type expected struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name         string
		args         args
		prepResponse func(fakeTestingT *assertfakes.FakeTestingT)
		testMock     func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT)
		expected     *expected
	}{
		{
			name: "success wantErr true",
			args: args{
				wantErr:    true,
				err:        errors.New("error"),
				msgAndArgs: []interface{}{"message", "hi"},
			},
			prepResponse: func(testingT *assertfakes.FakeTestingT) { /*noop*/ },
			testMock: func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT) {
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
			prepResponse: func(testingT *assertfakes.FakeTestingT) { /*noop*/ },
			testMock: func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT) {
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
			prepResponse: func(fakeTestingT *assertfakes.FakeTestingT) {

			},
			testMock: func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT) {
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
			prepResponse: func(fakeTestingT *assertfakes.FakeTestingT) {

			},
			testMock: func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *assertfakes.FakeTestingT) {
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
			r := vendorrequire.New(t)
			r.NotNil(test.prepResponse)
			r.NotNil(test.testMock)

			fakeTestingT := new(assertfakes.FakeTestingT)

			test.prepResponse(fakeTestingT)

			sut := assert.New(fakeTestingT)
			sut.WantError(test.args.wantErr, test.args.err, test.args.msgAndArgs)

			test.testMock(r, test.expected, fakeTestingT)
		})
	}
}
