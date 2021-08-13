package require_test

import (
	"errors"
	"testing"

	vendorrequire "github.com/stretchr/testify/require"

	"github.com/metrumresearchgroup/wrapt/require"
	"github.com/metrumresearchgroup/wrapt/require/requirefakes"
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
		prepResponse func(fakeTestingT *requirefakes.FakeTestingT)
		testMock     func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *requirefakes.FakeTestingT)
		expected     *expected
	}{
		{
			name: "success wantErr true",
			args: args{
				wantErr:    true,
				err:        errors.New("error"),
				msgAndArgs: []interface{}{"message %s", "hi"},
			},
			prepResponse: func(testingT *requirefakes.FakeTestingT) { /*noop*/ },
			testMock: func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *requirefakes.FakeTestingT) {
				r.Equal(0, fakeTestingT.ErrorfCallCount())
				r.Equal(0, fakeTestingT.FailNowCallCount())
			},
		},
		{
			name: "success wantErr false",
			args: args{
				wantErr:    false,
				err:        nil,
				msgAndArgs: []interface{}{"message %s", "hi"},
			},
			prepResponse: func(testingT *requirefakes.FakeTestingT) { /*noop*/ },
			testMock: func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *requirefakes.FakeTestingT) {
				r.Equal(0, fakeTestingT.ErrorfCallCount())
				r.Equal(0, fakeTestingT.FailNowCallCount())
			},
		},
		{
			name: "failure wantErr true",
			args: args{
				wantErr:    true,
				err:        nil,
				msgAndArgs: []interface{}{"message %s", "hi"},
			},
			prepResponse: func(fakeTestingT *requirefakes.FakeTestingT) {

			},
			testMock: func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *requirefakes.FakeTestingT) {
				r.Equal(1, fakeTestingT.ErrorfCallCount())
				r.Equal(1, fakeTestingT.FailNowCallCount())

				format, args := fakeTestingT.ErrorfArgsForCall(0)
				r.Equal(exp.format, format)
				r.Equal(exp.args, args)
			},
			expected: &expected{
				format: "\n%s",
				args:   []interface{}{"\tError Trace:\t\n\tError:      \tAn error is expected but got nil.\n\tMessages:   \t[message %s hi]\n"},
			},
		},
		{
			name: "failure wantErr false",
			args: args{
				wantErr:    false,
				err:        errors.New("new"),
				msgAndArgs: []interface{}{"message %s", "hi"},
			},
			prepResponse: func(fakeTestingT *requirefakes.FakeTestingT) {

			},
			testMock: func(r *vendorrequire.Assertions, exp *expected, fakeTestingT *requirefakes.FakeTestingT) {
				r.Equal(1, fakeTestingT.ErrorfCallCount())
				r.Equal(1, fakeTestingT.FailNowCallCount())

				format, args := fakeTestingT.ErrorfArgsForCall(0)
				r.Equal(exp.format, format)
				r.Equal(exp.args, args)
			},
			expected: &expected{
				format: "\n%s",
				args:   []interface{}{"\tError Trace:\t\n\tError:      \tReceived unexpected error:\n\t            \tnew\n\tMessages:   \t[message %s hi]\n"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := vendorrequire.New(t)
			r.NotNil(test.prepResponse)
			r.NotNil(test.testMock)

			fakeTestingT := new(requirefakes.FakeTestingT)

			test.prepResponse(fakeTestingT)

			sut := require.New(fakeTestingT)
			sut.WantError(test.args.wantErr, test.args.err, test.args.msgAndArgs)

			test.testMock(r, test.expected, fakeTestingT)
		})
	}
}
