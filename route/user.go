package route

import (
	"tes/auth"
	"tes/handler"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	route := r.Group("/user-api")
	{
		route.GET("user", auth.Auth, handler.GetUsers)
		route.POST("user", auth.Auth, handler.CreateUser)
		route.GET("user/:id", auth.Auth, handler.GetUserByID)
		route.PUT("user/:id", auth.Auth, handler.UpdateUser)
		route.DELETE("user/:id", auth.Auth, handler.DeleteUser)

		route.POST("login/:username/:password", handler.LoginAuthentication)
	}
	return r
}
