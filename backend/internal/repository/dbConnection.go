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

	err = CreateTables()

	if err != nil{
		return nil, err
	}

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

func CreateTables() error {
	createUsersTable := `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		CREATE TABLE IF NOT EXISTS users (
			id uuid DEFAULT uuid_generate_v4() NOT NULL,
			email text NOT NULL,
			password character varying(100) NOT NULL,
			first_name text NOT NULL,
			last_name text NOT NULL,
			is_active boolean DEFAULT true,
			is_staff boolean DEFAULT false,
			is_superuser boolean DEFAULT false,
			thumbnail text,
			created_at timestamp with time zone DEFAULT now() NOT NULL,
			
			CONSTRAINT users_pkey PRIMARY KEY (id),
			
			CONSTRAINT users_email_key UNIQUE (email)
		);

		CREATE INDEX IF NOT EXISTS users_id_email_is_active_indx ON users (id, email, is_active);
	`

	_, err := dbConn.SQL.Exec(createUsersTable)
	if err != nil{
		return err
	}

	createUsersProfileTable := `
		CREATE TABLE IF NOT EXISTS user_profile (
			id uuid DEFAULT uuid_generate_v4() NOT NULL,
			user_id uuid NOT NULL,
			phone_number text,
			birth_date date,
			
			CONSTRAINT user_profile_pkey PRIMARY KEY (id),
			
			CONSTRAINT user_profile_user_id_key UNIQUE (user_id),

			CONSTRAINT user_profile_user_id_fkey
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);

		CREATE INDEX IF NOT EXISTS users_detail_id_user_id ON user_profile (id, user_id);
	`

	_, err = dbConn.SQL.Exec(createUsersProfileTable)
	if err != nil{
		return err
	}

	return nil
}