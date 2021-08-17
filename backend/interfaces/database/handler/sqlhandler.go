package handler

import (
	"database/sql"
)

type SqlHandler interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	// get(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	//query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	//exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	//namedQuery(ctx context.Context, query string, namedParam map[string]interface{}) (*sqlx.Rows, error)
	//namedExec(ctx context.Context, query string, namedParam map[string]interface{}) (sql.Result, error)
	//rowsScan(ctx context.Context, rows *sqlx.Rows, typeI interface{}) (interface{}, error)
}
