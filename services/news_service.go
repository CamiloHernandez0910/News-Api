package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Cargar variables de entorno al iniciar
func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando .env:", err)
	}
}

// GetSources obtiene fuentes de noticias desde una API externa y permite filtrar por categoría, idioma y país.
//
// @Summary Obtener fuentes de noticias
// @Description Devuelve fuentes de noticias filtradas por categoría, idioma y país
// @Tags Fuentes
// @Accept json
// @Produce json
// @Param category query string false "Categoría de la fuente (ej: business, sports, technology)"
// @Param language query string false "Idioma de la fuente (ej: en, es, fr)"
// @Param country query string false "País de la fuente (ej: us, mx, ar)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /sources [get]
func GetSources(c *gin.Context) {
	category := c.Query("category")
	language := c.Query("language")
	country := c.Query("country")

	// 📌 Verificar si la API Key está disponible
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API Key no encontrada en .env"})
		return
	}

	// 📌 Construcción de la URL con filtros
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines/sources?apiKey=%s", apiKey)

	if category != "" {
		url += "&category=" + category
	}
	if language != "" {
		url += "&language=" + language
	}
	if country != "" {
		url += "&country=" + country
	}

	// 📌 Hacer la petición HTTP a la API de noticias
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener fuentes de noticias"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer respuesta"})
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar JSON"})
		return
	}

	c.JSON(http.StatusOK, data)
}
