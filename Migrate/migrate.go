package main

import (
	initializers "github.com/Shreyas-Prabhu/MySQLGO/Initializers"
	models "github.com/Shreyas-Prabhu/MySQLGO/Models"
)

func init() {
	initializers.LoadEnv()
	initializers.LoadDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
