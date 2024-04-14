package logic

import (
	"context"
	"server/internal/dal"
)

type TestLogic interface {
	Test(ctx context.Context)
}

type testLogic struct {
	testDal dal.TestDal
}

func NewTestLogic() TestLogic {
	return &testLogic{
		testDal: dal.NewTestDal(),
	}
}

func (t *testLogic) Test(ctx context.Context) {

}
