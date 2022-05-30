package api

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-resource-gin/model/common/response"
)

func Logout(c *gin.Context) {
	t := strings.Split(c.GetHeader("Authorization"), " ")
	if len(t) > 1 {
		token := t[1]
		resp, err := oauthService.Logout(token)
		if err == nil {
			response.OkWithData(resp, c)
			return
		}
	}
	response.Fail(c)
}
