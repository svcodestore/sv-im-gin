package system

import (
	"github.com/gin-gonic/gin"
	pb "github.com/svcodestore/sv-resource-gin/proto/auth"
	"github.com/svcodestore/sv-resource-gin/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type AuthService struct {
}

func (s *AuthService) AuthMenusWithApplicationIdAndUserId(applicationId, userId string) (menus gin.H, err error) {
	menus, err = utils.CallAuthRpcService(func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error) {
		c := pb.NewAuthClient(conn)

		r, e := c.GetAuthMenusWithApplicationIdAndUserId(ctx, &pb.GetAuthMenusWithApplicationIdAndUserIdRequest{
			ApplicationId: applicationId,
			UserId:        userId,
		})
		if e != nil {
			err = e
			return
		}
		reply = r.GetAuthMenus()
		return
	})

	return
}
