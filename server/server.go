package server

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func Start(db *sql.DB) {
	r := gin.Default()
	r.GET("/ping", HandlerGetAlgo(db))
	r.GET("/echo/:value", HandlerEcho(db))
	r.POST("/multiplicar", HandlerMultiplicacion())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func HandlerGetAlgo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var value string
		db.QueryRow("SELECT 'Hola!'").Scan(&value)
		c.JSON(200, gin.H{
			"message": value,
		})
	}
}

func HandlerEcho(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.Param("value")
		db.QueryRow("SELECT $1", value).Scan(&value)
		c.JSON(200, gin.H{
			"message": value,
		})
	}
}

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
			log.Println(err)
			return
		}

		resp := Response{
			Resultado: req.Multiplicando * req.Multiplicador,
		}
		c.JSON(200, resp)
	}
}
