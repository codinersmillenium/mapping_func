package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Database struct {
	DB     *sql.DB
	Driver string
}

func NewDatabase(driver, dsn string) (*Database, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("DB connected")
	return &Database{DB: db, Driver: driver}, nil
}

func (d *Database) Exec(q string, args ...any) (sql.Result, error) {
	return d.DB.Exec(d.Rebind(q), args...)
}

func (d *Database) Rebind(q string) string {
	if d.Driver != "postgres" {
		return q
	}

	i := 1
	out := ""
	for _, c := range q {
		if c == '?' {
			out += fmt.Sprintf("$%d", i)
			i++
		} else {
			out += string(c)
		}
	}
	return out
}
