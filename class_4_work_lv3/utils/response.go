package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}
func RespFail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"message": message,
	})
}
func LoginSuccess(c *gin.Context, message, tip1, tip2, tip3, tip4, tip5, tip6, tip7, tip8 string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
		"tip1":    tip1,
		"tip2":    tip2,
		"tip3":    tip3,
		"tip4":    tip4,
		"tip5":    tip5,
		"tip6":    tip6,
		"tip7":    tip7,
		"tip8":    tip8,
	})
}
func AllFriends(c *gin.Context, friend []string) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"friend": friend,
	})
}
