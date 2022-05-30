package service

import (
	"github.com/svcodestore/sv-resource-gin/service/system"
)

type Group struct {
	OauthService  system.OauthService
	AppService    system.AppService
	UserService   system.UserService
	AuthService   system.AuthService
}

var ServiceGroup = new(Group)
