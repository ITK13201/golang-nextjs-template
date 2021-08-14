package infrastructure

import (
	"database/sql"

	"github.com/ITK13201/go-next-test/backend/config"
	"github.com/ITK13201/go-next-test/backend/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(cfg *config.Config) *SqlHandler {
	conn, err := sql.Open("mysql", cfg.DatabaseUrl)
	if err != nil {
		panic(err.Error)
	}
	SqlHandler := new(SqlHandler)
	SqlHandler.Conn = conn
	return SqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}
