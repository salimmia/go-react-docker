package repository

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type DB struct{
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func ConnectDB(driver string, dsn string) (*DB, error){
	log.Println("Connecting to Database...")
	conn, err := NewDatabaseConnection(driver, dsn)
	if err != nil{
		log.Println("working...")
		return nil, err
	}

	conn.SetMaxOpenConns(maxOpenDbConn)
	conn.SetMaxIdleConns(maxIdleDbConn)
	conn.SetConnMaxIdleTime(maxDbLifetime)

	dbConn.SQL = conn

	log.Println("Connected to Database!")

	return dbConn, nil
}

func NewDatabaseConnection(driver string, dsn string) (*sql.DB, error){
	db, err := sql.Open(driver, dsn)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil{
		return nil, err
	}

	return db, nil
}