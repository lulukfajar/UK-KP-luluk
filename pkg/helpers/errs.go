package helpers

import "github.com/gin-gonic/gin"

func GenerateErrResponse(
	ctx *gin.Context,
	statusCode int,
	message string) {
	ctx.AbortWithStatusJSON(
		statusCode,
		map[string]any{
			"message": message,
		})
}
