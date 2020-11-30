package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/vue-go-crud/server/database"
	"github.com/hyperxpizza/vue-go-crud/server/handler"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8081"
	}

	user := "pizza"
	password := "Wojtekfoka1"
	dbname := "employeesdatabase"

	database.InitDB(user, password, dbname)

	router := gin.Default()

	//use default cors settings
	router.Use(cors.Default())

	//group api routes
	api := router.Group("/api")
	api.Use()
	{
		api.GET("/users", handler.GetAllEmployee)
		api.PUT("/update/:id", handler.UpdateEmployee)
		api.DELETE("/delete/:id", handler.DeleteUser)
		api.POST("/user", handler.AddUser)
	}

	router.Run(port)
}
