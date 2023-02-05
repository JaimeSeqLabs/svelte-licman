package sql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func WithDriverAndURL(driver string, url string) func(*orgSqlRepo) {
	return func(repo *orgSqlRepo) {
		db, err := sqlx.Connect(driver, url)
		if err != nil {
			panic(err)
		}
		repo.db = db
	}
}

func WithDatabase(db *sqlx.DB) func(*orgSqlRepo) {
	return func(org *orgSqlRepo) {
		org.db = db
	}
}

func WithSchemaMigration(schemas ...string) func(*orgSqlRepo) {
	return func(org *orgSqlRepo) {
		org.migrations = schemas
	}
}
