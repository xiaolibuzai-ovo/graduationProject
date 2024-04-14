package dal

import (
	"context"
	"database/sql"
	"server/internal/config"
)

type TestDal interface {
	Test(ctx context.Context)
}

type testDal struct {
	db *sql.DB
}

func NewTestDal() TestDal {
	return &testDal{
		db: config.GetDb(),
	}
}

func (t *testDal) Test(ctx context.Context) {

}
