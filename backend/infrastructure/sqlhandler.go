package infrastructure

import (
	"database/sql"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/ITK13201/golang-nextjs-template/backend/config"
)

type SqlHandler struct {
	Db *sqlx.DB
}

func NewSqlHandler(cfg *config.Config) handler.SqlHandler {
	db, err := sqlx.Connect("mysql", (*cfg).DatabaseUrl)
	if err != nil {
		panic(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Db = db

	return sqlHandler
}

func (handler *SqlHandler) NamedExec(query string, arg interface{}) (sql.Result, error) {
	res := SqlResult{}
	result, err := handler.Db.NamedExec(query, arg)
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
