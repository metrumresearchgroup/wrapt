package assert_test

import (
	"errors"
	"testing"

	"github.com/metrumresearchgroup/wrapt/assert"
	"github.com/metrumresearchgroup/wrapt/testingt/mock"
)

func TestAssertions_WantError(t *testing.T) {
	type args struct {
		wantErr    bool
		err        error
		msgAndArgs []interface{}
	}
	tests := []struct {
		name     string
		args     args
		expected *mock.TestingT
	}{
		{
			name: "success",
			args: args{
				wantErr:    true,
				err:        errors.New("error"),
				msgAndArgs: []interface{}{"message %s", "hi"},
			},
			expected: &mock.TestingT{
				HelperCalled: true,
			},
		},
		{
			name: "failure",
			args: args{
				wantErr:    true,
				err:        nil,
				msgAndArgs: []interface{}{"message %s", "hi"},
			},
			expected: &mock.TestingT{
				Format:        "\n%s",
				Args:          []interface{}{"\tError Trace:\t\n\tError:      \tAn error is expected but got nil.\n\tMessages:   \t[message %s hi]\n"},
				ErrorfCalled:  true,
				FailNowCalled: false,
				HelperCalled:  true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tt2 := &mock.TestingT{}
			req := assert.New(tt2)

			req.WantError(test.args.wantErr, test.args.err, test.args.msgAndArgs)

			a := assert.New(t)
			a.Equal(test.expected, tt2)
		})
	}
}
