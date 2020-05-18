ppackage database

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5431
	user     = "postgres"
	password = "cph20429"
	dbname	 = "registerdb"
)

func InitPostgresql() {
	psqlInfo := fmt.Sprintf("user=%s "+"password=%s dbname=%s sslmode=disable",user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row { 
 	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
 	return db.Query(sql)
}