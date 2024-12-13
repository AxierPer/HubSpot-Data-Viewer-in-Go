package main

import (
	rt "api_hubspot_go/src/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	ID string `uri:"id"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./static/templates/*")
	r.Static("/static/", "./static/styles/")

	r.GET("/", func(c *gin.Context) {
		response := rt.GetClients()
		c.HTML(http.StatusOK, "general.html", gin.H{
			"Items": response,
		})
	})

	r.GET("/:id", func(c *gin.Context) {
		var data Data
		if err := c.ShouldBindUri(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		response := rt.GetClientId(data.ID)
		fmt.Printf("\n tipo: %T \n", response)
		c.HTML(http.StatusOK, "getId.html", gin.H{
			"title":    "Main",
			"name":     response["firstname"],
			"lastname": response["lastname"],
			"id":       response["hs_object_id"],
			"email":    response["email"],
		})

	})

	r.Run() // Listen and server on 0.0.0.0:8080
}
