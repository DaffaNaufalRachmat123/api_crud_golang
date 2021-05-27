package config

import (
	"strconv"
	"github.com/gin-gonic/gin"
)
const PaginationLimit int = 10
const PaginationOffset int = 10

func Page(c *gin.Context) int {
	page := c.Query("p")
	var pages int 
	pages , err := strconv.Atoi(page)
	if err != nil {
		pages = 0
	}
	return pages
}