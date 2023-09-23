package sample

import (
	"github.com/gin-gonic/gin"
)

func DbInsert(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "dbinsert",
	})

}
