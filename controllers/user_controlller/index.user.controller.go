package user_controlller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {
	isValidated := true
	if !isValidated {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "error is not validated",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello user",
	})
}
