package infrastructure

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ITK13201/golang-nextjs-template/backend/config"
	"github.com/ITK13201/golang-nextjs-template/backend/domain"
	"github.com/ITK13201/golang-nextjs-template/backend/interfaces/database"
)

type SqlHandler struct {
	DbMap *gorp.DbMap
}

func NewSqlHandler(cfg *config.Config) database.SqlHandler {
	db, err := sql.Open("mysql", (*cfg).DatabaseUrl)
	if err != nil {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	// Mapping structures to tables
	dbmap.AddTableWithName(domain.User{}, "users").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		panic(err)
	}

	return dbmap
}
