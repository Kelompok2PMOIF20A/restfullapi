package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerHelloWorld(c *gin.Context) {

	c.JSON(http.StatusOK, "Hello World!")

}
