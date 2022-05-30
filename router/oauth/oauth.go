package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-resource-gin/api/app"

	"github.com/svcodestore/sv-resource-gin/api/oauth"
)

type OAuthRoutes struct {
}

func (*OAuthRoutes) Init(r *gin.RouterGroup) {
	appG := r.Group("application")
	appG.GET("/current-application", app.GetCurrentApp)

	oauthG := r.Group("oauth2.0")
	oauthG.POST("/token", oauth.GetAccessToken)
}
