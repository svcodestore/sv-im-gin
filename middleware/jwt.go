package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-resource-gin/model/common/response"
	"github.com/svcodestore/sv-resource-gin/service"
)

var oauthService = service.ServiceGroup.OauthService

func JWTCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqPath := c.FullPath()
		if reqPath == "/api/oauth2.0/token" || reqPath == "/api/logout" {
			c.Next()
			return
		}

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.UnAuth(c)
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer") {
			t := strings.Split(token, " ")
			if len(t) != 2 || t[1] == "" {
				response.UnAuth(c)
				c.Abort()
				return
			}
			token = t[1]
		}

		isLogin, claims, _ := oauthService.IsUserLogin(token)

		if !isLogin {
			response.UnAuth(c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
