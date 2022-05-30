package user

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-resource-gin/model/common/response"
	"github.com/svcodestore/sv-resource-gin/utils"
)

func GetCurrentUser(c *gin.Context) {
	id := utils.GetUserID(c)
	user, _ := userService.UserWithId(id)
	response.OkWithData(user, c)
}

