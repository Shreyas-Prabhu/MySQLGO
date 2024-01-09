package main

import (
	"fmt"

	controllers "github.com/Shreyas-Prabhu/MySQLGO/Controllers"
	initializers "github.com/Shreyas-Prabhu/MySQLGO/Initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.LoadDB()
}

func main() {
	fmt.Println("Hello...sWelcome to GORM, MYSQL, GIn, GO, CRUD")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	r.GET("/user/:id", controllers.GetControllerbyId)
	r.GET("/user", controllers.GetAllController)
	r.POST("/user", controllers.PostController)
	r.PUT("/user/:id", controllers.UpdateController)
	r.DELETE("/user/:id", controllers.DeleteController)
	r.Run()

}
