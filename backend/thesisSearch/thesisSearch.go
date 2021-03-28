package thesisSearch

import (
	"backend/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

// GetThesisList 查找相应文章
func GetThesisList(c *gin.Context) {
	source := c.Query("source")
	year := c.Query("year")
	keyword := c.Query("keyword")

	keyword = strings.ReplaceAll(keyword, "[^A-Za-z]", "%")  //problem:正则表达式
	keyword = "%" + keyword + "%"

	var selectStr string = "select * from analyzed_thesis where "
	if source != "" {
		selectStr += "Source = '" + source + "' "
	}
	if year != "" {
		if source != "" {
			selectStr += "and "
		}
		selectStr += "Year = '" + year + "' "
	}
	if keyword != "" {
		if !(source == "" && year == "") {
			selectStr += "and "
		}
		selectStr += "Keyword like '" + keyword + "' "
	}

	if source == "" && year == "" && keyword == "" {
		length := len(selectStr) - 6
		selectStr = selectStr[0:length]
	}
	selectStr += " ;"

	rows, _ := database.DB.Raw(selectStr).Rows()
	var ID int
	var Source string
	var Year int
	var Title string
	var Author string
	var Keyword string
	var Abstract string
	for rows.Next() {
		_ = rows.Scan(&ID, &Source, &Year, &Title, &Author, &Keyword, &Abstract)
	}

	c.JSON(http.StatusOK, gin.H{"Title":Title})
}
