package system

import (
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-resource-gin/model/system/request"
	pb "github.com/svcodestore/sv-resource-gin/proto/oauth"
	"github.com/svcodestore/sv-resource-gin/utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type OauthService struct {
}

func (s *OauthService) GetAccessToken(grantType, clientId, clientSecret, code, redirectUri string) (resp gin.H, err error) {
	resp, err = utils.CallSsoRpcService(func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error) {
		c := pb.NewOauthClient(conn)

		r, e := c.GetOauthCode(ctx, &pb.GetOauthCodeRequest{
			GrantType:    grantType,
			ClientId:     clientId,
			ClientSecret: clientSecret,
			Code:         code,
			RedirectUri:  redirectUri,
		})
		if e != nil {
			log.Fatalf("could not get user: %v", e)
			return
		}

		reply = r.GetOauthInfo()
		return
	})
	return
}

func (s *OauthService) IsUserLogin(accessToken string) (isLogin bool, claims *request.CustomClaims, err error) {
	result, err := utils.CallSsoRpcService(func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error) {
		c := pb.NewOauthClient(conn)

		r, e := c.IsUserLogin(ctx, &pb.IsUserLoginRequest{AccessToken: accessToken})
		if e != nil {
			log.Fatalf("could not get user: %v", e)
			return
		}

		reply = r.GetIsUserLoginResult()
		return
	})
	if err != nil {
		return
	}
	code := int(result["code"].(float64))
	data := result["data"].(map[string]interface{})
	message := result["message"].(string)
	if code != 0 {
		err = errors.New(message)
		return
	}

	isLogin = data["isLogin"].(bool)
	b, err := json.Marshal(data["claims"])
	if err != nil {
		return
	}
	json.Unmarshal(b, &claims)

	return
}

func (s *OauthService) Logout(accessToken string) (resp gin.H, err error) {
	resp, err = utils.CallSsoRpcService(func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error) {
		c := pb.NewOauthClient(conn)

		r, e := c.Logout(ctx, &pb.LogoutRequest{AccessToken: accessToken})
		if e != nil {
			log.Fatalf("could not get user: %v", e)
			return
		}

		reply = r.GetLogoutResult()
		return
	})
	return
}
