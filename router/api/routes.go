package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-resource-gin/api"
	"github.com/svcodestore/sv-resource-gin/api/auth"
	"github.com/svcodestore/sv-resource-gin/api/user"
)

type Routes struct {
}

func (*Routes) Init(r *gin.RouterGroup) {
	r.POST("logout", api.Logout)

	userG := r.Group("user")
	userG.GET("current-user", user.GetCurrentUser)

	authG := r.Group("authorization")
	authG.GET("/user-menus", auth.GetAuthMenus)
}
