package presenter

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/teran/echo-grpc-server/presenter/proto"
)

type EchoHandlers interface {
	proto.EchoServiceServer

	Register(*grpc.Server)
}

type echoHandlers struct {
	proto.UnimplementedEchoServiceServer
}

func NewEchoHandlers() EchoHandlers {
	return &echoHandlers{}
}

func (h *echoHandlers) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	msg := req.GetMessage()

	log.WithFields(log.Fields{
		"message": msg,
	}).Infof("message received")

	return proto.NewEchoResponse(msg)
}

func (h *echoHandlers) Register(srv *grpc.Server) {
	proto.RegisterEchoServiceServer(srv, h)
}
