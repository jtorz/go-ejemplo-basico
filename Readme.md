# Go Ejemplo Básico

Este proyecto es un ejemplo básico de un servidor HTTP en Go, usando el framework [Gin-Gonic](https://github.com/gin-gonic/gin)
y base de datos postgres.

## Tutorial

### 1. Descargar el proyecto

Puedes clonar el proyecto o descargalo como zip desde [aqui](https://github.com/jtorz/go-ejemplo-basico/archive/refs/heads/main.zip).

### 2. Ejecutar
Para ejecutar este ejemplo se necesita registrar
la cadena de conexión a una base de datos postgres en la variable de entorno
`EJEMPLO_DB_CONNECTION`

La cadena debe tener el siguiente formato:
```
host=localhost port=5432 user=postgres dbname=dbname password=password sslmode=disable
```

Para correrlo al igual que cualquier aplicación de Go
se puede ejecutar el comando desde la raíz del proyecto:

```
go run .
```

Esto iniciara un servidor HTTP en el puerto 8080 con tres rutas:

* `GET /ping`: regresa un json con el mensaje Hola!
* `GET /echo/algun-valor`: regresa un json con el mensaje algun-valor.
* `POST /multiplicar`: multiplica dos numeros.

## NOTA
Si se utiliza VS Code se puede copiar el archivo `dev.env.example` a `dev.env` y ejecutar en modo debug con F5.

### Probar

Para probar los endpoint `GET` pueden entrar a las rutas:

* [localhost:8080/ping](http://localhost:8080/ping)
* [localhost:8080/echo/puedes cambiar este valor](http://localhost:8080/echo/puedes%20cambiar%20este%20valor)

Para pobar el endpoint `POST /multiplicar` se puede hacer desde postman, o intalando la extension de VS Code [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) y ejecutar las llamadas desde el archivo `rest_client/rest_client.rest`

## Organizacion del proyecto

Entre todos los archivos los dos mas importantes son:

1. server/server.go
2. main.go

### main.go

```go
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

```
#### imports

En la primera parte se tienen los import de los paquetes utilizados en este archivo.

* `"database/sql"`: Se utiliza para conectarse a la base de datos.
* `"fmt"`: Se utiliza para imprimir en la consola.
* `"os"`: Se utiliza para obtener valores desde las variables de entorno.
* `"github.com/jtorz/go-ejemplo-basico/server"`: Paquete de implementasion del ejemplo.
* `_ "github.com/lib/pq"`: Driver de conexion a postgres. (Este paquete no se manda a llamar directamente pero es necesario para conectarse postgres, por eso se usa el guion bajo)

#### func ConnectPostgres

La funcion connect postgres se encarga de conectarse a la base de datos y realizar un ping para probar la conexión.

### func main
La funcion main solo imprime un valor, en consola, manda a llamar a la funcion ConnectPostgres para conectarse a la base de datos y finalmente llama a la funcion Start del paquete `"github.com/jtorz/go-ejemplo-basico/server"` en el cual se configura e inicia el servidor.

## server/server.go

```go
func Start(db *sql.DB) {
	r := gin.Default()
	r.GET("/ping", HandlerGetAlgo(db))
	r.GET("/echo/:value", HandlerEcho(db))
	r.POST("/multiplicar", HandlerMultiplicacion())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

En la funcion Start se configura el servidor. Esta funcion recibe la base de datos a utilizar.

Aqui se puede observar que existen 3 endpoints y las funciones que los manejan .

Cada una de las tres funciones regresan una funcion del tipo `gin.HandlerFunc`. Esta funcion es la que se manda a llamar cada vez que se hace una llamada desde el cliente.

Esto permite aislar informacion entre los distintos handlers, como se puede ver en la funcion `HandlerMultiplicacion` en la cual se declaran las estructuras de las peticiones y las respuestas.


Si en algun momento ocurre un error en el handler, se puede observar el tratamiento del error.