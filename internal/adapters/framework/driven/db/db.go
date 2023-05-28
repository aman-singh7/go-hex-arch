package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type DbAdapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*DbAdapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v\n", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v\n", err)
	}

	return &DbAdapter{
		db: db,
	}, nil
}

func (d *DbAdapter) CloseDbConnection() {
	err := d.db.Close()
	fmt.Println("db is closing")
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (d *DbAdapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()

	if err != nil {
		return err
	}

	_, err = d.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
