package system

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/svcodestore/sv-resource-gin/proto/user"
	"github.com/svcodestore/sv-resource-gin/utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type UserService struct {
}

func (s *UserService) UserWithId(id string) (user gin.H, err error) {
	user, err = utils.CallSsoRpcService(func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error) {
		c := pb.NewUserClient(conn)

		r, e := c.GetUserById(ctx, &pb.GetUserByIdRequest{
			Id: id,
		})
		if e != nil {
			err = e
			return
		}
		reply = r.GetUser()
		return
	})
	return
}

func (s *UserService) UsersWithApplicationId(applicationId string) (users gin.H, err error) {
	users, err = utils.CallSsoRpcService(func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error) {
		c := pb.NewUserClient(conn)

		r, e := c.GetAvailableUsersByApplicationId(ctx, &pb.GetAvailableUsersByApplicationIdRequest{
			ApplicationId: applicationId,
		})
		if e != nil {
			err = e
			return
		}
		reply = r.GetUsers()
		return
	})
	return
}