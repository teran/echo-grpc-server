package presenter

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/teran/echo-grpc-server/presenter/proto"
)

type RemoteEchoHandlers interface {
	proto.RemoteEchoServiceServer

	Register(*grpc.Server)
}

type remoteEchoHandlers struct {
	proto.UnimplementedRemoteEchoServiceServer
}

func NewRemoteEchoHandlers() RemoteEchoHandlers {
	return &remoteEchoHandlers{}
}

func (h *remoteEchoHandlers) RemoteEcho(ctx context.Context, req *proto.RemoteEchoRequest) (*proto.RemoteEchoResponse, error) {
	var (
		remote = req.GetRemote()
		msg    = req.GetMessage()
	)

	log.WithFields(log.Fields{
		"remote":  remote,
		"message": msg,
	}).Infof("message received")

	dial, err := grpc.Dial(remote, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	cli := proto.NewEchoServiceClient(dial)
	resp, err := cli.Echo(ctx, &proto.EchoRequest{
		Message: msg,
	})
	if err != nil {
		return nil, err
	}

	return proto.NewRemoteEchoResponse(resp.GetMessage())
}

func (h *remoteEchoHandlers) Register(srv *grpc.Server) {
	proto.RegisterRemoteEchoServiceServer(srv, h)
}
