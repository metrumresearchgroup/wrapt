package testingt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/metrumresearchgroup/wrapt/testingt/mock"
)

func TestMock(t *testing.T) {
	mtt := mock.TestingT{}
	req := require.New(t)

	req.Equal(false, mtt.ErrorfCalled)
	req.Equal(false, mtt.FailNowCalled)

	mtt.Errorf("a %s, %d, %v", "b", 1, errors.New("new err"))

	req.Equal(true, mtt.ErrorfCalled)
	req.Equal("a %s, %d, %v", mtt.Format)
	req.Equal([]interface{}{"b", 1, errors.New("new err")}, mtt.Args)

	mtt.FailNow()

	req.Equal(mtt.FailNowCalled, true)
}
