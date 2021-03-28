package thesisSearch

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetThesisList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name":"list"})
}
