/*
* File              : connect.go
* Author            : Eliazar
* Creation date     : 10/03/2024
* Last modified by  : Eliazar
* Last modified date: 10/03/2024
* Description       : This file contains functions to connect to a MySQL database.
 */

package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"database/sql"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dns := fmt.Sprintf("%v:%v@tcp(%v:%s)/%v",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("The connection to the MySQL database is successful")

	return db, nil
}
