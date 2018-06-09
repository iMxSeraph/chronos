package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(context *gin.Context) {
	context.String(http.StatusOK, "test")
}
