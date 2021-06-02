# Go Ejemplo Básico

Este proyecto es un ejemplo básico de un servidor HTTP en Go, usando el framework [Gin-Gonic](https://github.com/gin-gonic/gin)
y base de datos postgres.

## Tutorial

### 1. Descargar el proyecto

Puedes descargar el proyecto de diversas formas una desde github.

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

