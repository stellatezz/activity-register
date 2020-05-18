package main

import (
	"github.com/ivancc666/tkp-register-go/database"
    "github.com/ivancc666/tkp-register-go/routers"
)

func main() {
	database.InitPostgresql()
	router := routers.InitRouter()
    
    router.Static("/static", "./static")
    router.Run(":8000")
}