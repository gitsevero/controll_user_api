package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"main/Routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	Routes.Routes(&router.RouterGroup)
	err = router.Run(":8080")
	if err != nil {
		return
	}

}
