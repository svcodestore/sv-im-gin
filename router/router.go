package router

import (
	"github.com/svcodestore/sv-resource-gin/router/api"
	"github.com/svcodestore/sv-resource-gin/router/oauth"
)

type Group struct {
	OAuth oauth.OAuthRoutes
	Api   api.Routes
}

var RouterGroup = new(Group)
