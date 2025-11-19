package middleware

import (
	"net/http"

	"github.com/abdallahelassal/Store/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc{
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		utils.ErrorResponse(c,http.StatusInternalServerError,"internal server error","")
		c.Abort()
	})
}