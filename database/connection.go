package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/bloiseleo/chorumecon/env"
	_ "github.com/go-sql-driver/mysql"
)

var (
	conn     *sql.DB
	connOnce sync.Once
)

func createConnectionString() string {
	user := env.Env("DB_USER", "application")
	pass := env.Env("DB_PASS", "application")
	db := env.Env("DB_NAME", "chorume_coins")
	host := env.Env("DB_HOST", "localhost")
	port := env.Env("DB_PORT", "3306")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, db)
}

func realConnect() {
	connectionString := createConnectionString()
	c, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	conn = c
}

func Connect() *sql.DB {
	connOnce.Do(realConnect)
	return conn
}
