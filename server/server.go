package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Start configura e inicia el servidor
func Start(db *sql.DB) {
	r := gin.Default()
	r.GET("/ping", HandlerGetAlgo(db))
	r.GET("/echo/:value", HandlerEcho(db))
	r.POST("/multiplicar", HandlerMultiplicacion())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// HandlerGetAlgo crea un mensaje consultado desde la base de datos.
func HandlerGetAlgo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var value string
		err := db.QueryRow("SELECT 'Hola!'").Scan(&value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Úps, esto no deberia pasar",
				"err":     err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": value,
		})
	}
}

// HandlerEcho crea un mensaje con el valor del parametro :value.
// Haciendo eco desde la base de datos
func HandlerEcho(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		urlParam := c.Param("value")

		var value string
		err := db.QueryRow("SELECT $1", urlParam).Scan(&value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Úps, esto no deberia pasar",
				"err":     err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": value,
		})
	}
}

// HandlerMultiplicacion realiza la multiplicación de dos numeros..
func HandlerMultiplicacion() gin.HandlerFunc {
	type Request struct {
		Multiplicando float64
		Multiplicador float64
	}
	type Response struct {
		Resultado float64
	}

	return func(c *gin.Context) {
		req := Request{}
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Úps, esto no deberia pasar",
				"err":     err.Error(),
			})
			return
		}

		resp := Response{
			Resultado: req.Multiplicando * req.Multiplicador,
		}
		c.JSON(200, resp)
	}
}
