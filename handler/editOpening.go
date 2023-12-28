package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}
