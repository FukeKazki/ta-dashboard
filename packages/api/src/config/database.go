package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	Conn *sql.DB
}

func dsn() string {
	c := mysqlDriver.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		DBName:               os.Getenv("DB_NAME"),
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Net:                  "tcp",
		ParseTime:            true,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}
	c.Params = make(map[string]string, len(c.Params))
	c.Params["charset"] = "utf8mb4"
	c.Params["collation"] = "utf8mb4_general_ci"
	c.Params["time_zone"] = fmt.Sprintf("'%s'", c.Loc.String())
	return c.FormatDSN()
}

func NewClient() (*DBClient, error) {
	conn, err := sql.Open("mysql", dsn())
	if err != nil {
		return nil, err
	}
	return &DBClient{Conn: conn}, nil
}
