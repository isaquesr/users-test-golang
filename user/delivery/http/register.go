package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isaquesr/users-test-golang/user"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc user.UseCase) {
	h := NewHandler(uc)

	users := router.Group("/users")
	{
		users.POST("", h.Create)
		users.GET("", h.Get)
		users.DELETE("", h.Delete)
	}
}
