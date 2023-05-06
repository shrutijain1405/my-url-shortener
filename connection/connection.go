package connection

import (
    "database/sql"
    "github.com/go-sql-driver/mysql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var (
    Db *sql.DB
)

func MakeDbConnection() {
    cfg := mysql.Config{
        User:   "root",
        Passwd: "Shruti@123",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "url-shortner",
    }
    // Get a database handle.
    var err error
    Db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := Db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    log.Println("Connected!")
}
