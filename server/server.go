package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/vue-go-crud/server/database"
	"github.com/hyperxpizza/vue-go-crud/server/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	database.InitDB(user, password, dbname)

	router := gin.Default()

	//use default cors settings
	router.Use(cors.Default())

	//group api routes
	api := router.Group("/api")
	api.Use()
	{
		api.GET("/users", handler.GetAllEmployee)
		api.PUT("/update", handler.UpdateEmployee)
		api.DELETE("/delete/:id", handler.DeleteUser)
		api.POST("/user", handler.AddUser)
	}

	router.Run(port)
}
