package main

import (
	"fmt"
	"log"

	_ "api-noticias/docs"
	"api-noticias/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Noticias
// @version 1.0
// @description API que proporciona fuentes de noticias con filtros por categoría, idioma y país.
// @host localhost:8080
// @BasePath /api
func main() {
	r := gin.Default()

	// 📌 Endpoint de Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 📌 Grupo de rutas con prefijo "/api"
	api := r.Group("/api")
	{
		api.GET("/sources", services.GetSources) // ✅ Nueva ruta
	}

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
