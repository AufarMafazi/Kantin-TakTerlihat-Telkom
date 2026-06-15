package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func InitDB() *sql.DB {
    dsn := "root:@tcp(127.0.0.1:3306)/tugasbesariae"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    return db
}