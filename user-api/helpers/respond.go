package helpers

import (
	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, status int, message string, err error) {
	resp := gin.H{"erro": message}
	if err != nil {
		resp["detalhes"] = err.Error()
	}
	c.JSON(status, resp)
}
