package main

import (
	"log"
	"url_shortener/cmd/api/handler"
	"url_shortener/infrastructure/database"
	"url_shortener/repository"
	"url_shortener/service"

	"github.com/gin-gonic/gin"
)

func main() {

	sqlite, errSqlite := database.NewSqlite()

	if errSqlite != nil {
		log.Fatal(errSqlite)
	}

	urlRepository := repository.NewUrlDBSqlite(sqlite.DB)
	urlService := service.NewURLService(urlRepository)

	urlHandler := handler.URLHandler{Service: urlService}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome guri",
		})
	})
	r.GET("/:hash", urlHandler.Redirect)
	r.POST("/", urlHandler.Save)
	r.Run(":1234")
}
