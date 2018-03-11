package main

import (
	"rest-api-server/app"

	"github.com/gin-gonic/gin"

	"fmt"
)

func main() {
	fmt.Println("Starting...")

	router := SetupRouter()
	router.Run(":8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/instructions", app.GetInstructions)
		v1.GET("/instructions/:id", app.GetInstruction)
		v1.POST("/instructions", app.PostInstruction)
		v1.PUT("/instructions/:id", app.UpdateInstruction)
		v1.DELETE("/instructions/:id", app.DeleteInstruction)
	}

	return router
}
