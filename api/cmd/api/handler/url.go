package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"url_shortener/domain/entity"
	"url_shortener/lib/logger"
	"url_shortener/service"
)

type URLHandler struct {
	Service *service.URLService
}

func (u *URLHandler) Save(c *gin.Context)  {
	var url entity.URL

	if err := c.BindJSON(&url) ; err != nil {
		logger.Error(fmt.Sprintf("an error occurred while parsing json: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "an error occurred while parsing json"})
		return
	}

	result, err := u.Service.Save(url)
	if err != nil {
		logger.Error(fmt.Sprintf("an error occurred saving the url: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "an error occurred saving the url"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u *URLHandler) Find(c *gin.Context)  {
	hash := c.Param("hash")

	url, err := u.Service.Find(hash)

	if err != nil {
		logger.Error(fmt.Sprintf("an error occurred fetching url with this hash: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred fetching url with this hash"})
		return
	}

	if url.ID == 0 {
		logger.Error("not found url: %s")
		c.JSON(http.StatusNotFound, gin.H{})
		return 
	}

	c.JSON(http.StatusOK, url)
}