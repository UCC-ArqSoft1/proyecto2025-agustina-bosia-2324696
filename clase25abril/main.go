package main

import (
	"clase25abril/handlers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Ejemplo de API REST con Gin")
	router := gin.Default()

	router.GET("/HelloWorld", handlers.HelloWorld)

	err := router.Run(":80")
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}

	fmt.Println("Servidor corriendo en http://localhost:80")
}
