package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"	
    "learn-go/models"
	"gorm.io/gorm"
    "github.com/google/uuid"
)

type ClientRequest struct {
    Message string `json:"message"`
}

func Handler(db *gorm.DB) {


    router := gin.Default()

    router.GET("/api/v1/healthcheck", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"messaage": "service running"})
    })
    
	router.GET("/api/v1/urls/:url", func(c *gin.Context) {
        originalURL := c.Param("url")
        var url models.URL

        result := db.Where("shorten_url = ?", originalURL).First(&url)
        if result.Error != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"shorten_url": url.OriginalURL})
    })

    router.POST("/api/v1/urls", func(c *gin.Context) {
        var req ClientRequest
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }
        newURL := CreateNewURL(req.Message)
		url := models.URL{
			OriginalURL: req.Message, 
			ShortenURL: newURL,
		}
		result := db.Create(&url) 
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Database Error"})
		}else {
			c.JSON(http.StatusOK, gin.H{"message": newURL})
		}
    })

    router.Run("0.0.0.0:8080")
}

func CreateNewURL(oldURL string) string {
    out := uuid.Must(uuid.NewRandom()).String()[:8]
	return out
}
