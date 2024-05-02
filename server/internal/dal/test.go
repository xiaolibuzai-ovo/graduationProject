package dal

import (
	"context"
)

type TestDal interface {
	Test(ctx context.Context)
}

type testDal struct {
}

func NewTestDal() TestDal {
	return &testDal{}
}

func (t *testDal) Test(ctx context.Context) {

}
