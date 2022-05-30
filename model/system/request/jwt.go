package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserId   string    `json:"userId"`
	UUID     uuid.UUID `json:"uuid"`
	LoginId  string    `json:"loginId"`
	Username string    `json:"username"`
	NickName string    `json:"nickName"`
	ClientId string    `json:"clientId"`
}
