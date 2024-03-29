package main

import (
	"database/sql"
	"fmt"
	"os"

	// postgres driver
	"github.com/jtorz/go-ejemplo-basico/server"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Welcome!!")
	db, err := connectPostgres(os.Getenv("EJEMPLO_DB_CONNECTION"))
	if err != nil {
		fmt.Println("Ha ocurrido un error al conectarse a la base de datos: ", err)
		os.Exit(1)
	}
	server.Start(db)
}

func connectPostgres(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
