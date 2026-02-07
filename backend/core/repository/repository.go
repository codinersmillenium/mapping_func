package repository

import "database/sql"

type DB interface {
	Exec(query string, args ...any) (sql.Result, error)
	Rebind(string) string
}
