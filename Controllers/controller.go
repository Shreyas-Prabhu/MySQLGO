package controllers

import (
	"fmt"
	"log"

	initializers "github.com/Shreyas-Prabhu/MySQLGO/Initializers"
	models "github.com/Shreyas-Prabhu/MySQLGO/Models"
	"github.com/gin-gonic/gin"
)

func PostController(c *gin.Context) {
	var body struct {
		Name    string
		Company string
	}
	c.Bind(&body)
	usr := models.User{Name: body.Name, Company: body.Company}
	res := initializers.DB.Create(&usr)
	if res.Error != nil {
		log.Fatal("Cannot push the employee")
	}
	fmt.Println(res)
	c.JSON(200, gin.H{
		"message": "Inserted",
	})
}

func GetControllerbyId(c *gin.Context) {
	id := c.Param("id")
	var usr models.User
	_ = initializers.DB.First(&usr, id)
	c.JSON(200, gin.H{
		"Got details": usr,
	})

}

func GetAllController(c *gin.Context) {
	var usr []models.User
	result := initializers.DB.Find(&usr)
	if result.Error != nil {
		log.Fatal("Cannot push the employee")
	}
	c.JSON(200, gin.H{
		"All Employee": usr,
	})
}

func UpdateController(c *gin.Context) {

	id := c.Param("id")
	var emp models.User
	initializers.DB.First(&emp, id)

	var body struct {
		Name    string
		Company string
	}

	c.Bind(&body)

	initializers.DB.Model(&emp).Updates(models.User{Name: body.Name, Company: body.Company})

	c.JSON(200, gin.H{
		"Updated Employee": emp,
	})
}

func DeleteController(c *gin.Context) {

	id := c.Param("id")
	initializers.DB.Delete(&models.User{}, id)

	c.JSON(200, gin.H{
		"message": "deleted",
	})
}
