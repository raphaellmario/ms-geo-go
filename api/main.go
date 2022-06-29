package main

import (
	"tecban"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/all", getAllTerminais)
	router.GET("/find/:id", findTerminal)

	router.Run("localhost:8080")
}

func getAllTerminais(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	resposta := tecban.GetAll(latitude, longitude)
	c.IndentedJSON(http.StatusOK, resposta)
}

func findTerminal(c *gin.Context) {
	id := c.Param("id")
	resposta := tecban.Find(id)
	c.IndentedJSON(http.StatusOK, resposta)
}
