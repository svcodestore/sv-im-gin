package system

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/svcodestore/sv-resource-gin/proto/application"
	"github.com/svcodestore/sv-resource-gin/utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type AppService struct {
}

func (s *AppService) AppClientSecretWithClientId(clientId string) (clientSecret string, err error) {
	clientSecret, err = utils.CallSsoRpcServiceStringReply(func(conn *grpc.ClientConn, ctx context.Context) (reply string, e error) {
		c := pb.NewApplicationClient(conn)

		r, e := c.GetApplicationSecretByClientId(ctx, &pb.GetApplicationSecretByClientIdRequest{
			ClientId: clientId,
		})
		if e != nil {
			return
		}
		reply = r.GetClientSecret()
		return
	})
	return
}

func (s *AppService) AppWithId(id string) (app gin.H, err error) {
	app, err = utils.CallSsoRpcService(func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error) {
		c := pb.NewApplicationClient(conn)

		r, e := c.GetApplicationById(ctx, &pb.GetApplicationByIdRequest{
			Id: id,
		})
		if e != nil {
			log.Fatalf("could not get user: %v", e)
			return
		}
		reply = r.GetApplication()
		return
	})
	return
}

func (s *AppService) Apps() (app gin.H, err error) {
	app, err = utils.CallSsoRpcService(func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error) {
		c := pb.NewApplicationClient(conn)

		r, e := c.GetAvailableApplications(ctx, &pb.GetAvailableApplicationsRequest{})
		if e != nil {
			log.Fatalf("could not get user: %v", e)
			return
		}
		reply = r.GetApplications()
		return
	})
	return
}
