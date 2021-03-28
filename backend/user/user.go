package user

import (
	_ "backend/database"
	_ "database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strings"
)


type AnalyzedThesis struct {
	ID int
	Source string
	Year int
	Title string
	Author string
	Keyword string
	Abstract string
}

func UserLogin(c *gin.Context){
	kw := c.Query("kw")
	year := c.Query("year")
	c.JSON(http.StatusOK, gin.H{"name":kw, "year":year})
}

func UserRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name":"list"})
}

func GetThesisList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name":"list"})
}

