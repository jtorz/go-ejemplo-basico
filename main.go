package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// postgres driver
	"github.com/jtorz/go-ejemplo-basico/server"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Welcome!!")
	db, err := ConnectPostgres(os.Getenv("EJEMPLO_DB_CONNECTION"))
	if err != nil {
		log.Println("Ha ocurrido un error al conectarse a la base de datos: ", err)
		os.Exit(1)
	}
	server.Start(db)
}

func ConnectPostgres(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	// Query para probar la conexion.
	var one int
	err = db.QueryRow("SELECT 1").Scan(&one)
	if err != nil {
		return nil, err
	}
	return db, nil
}
