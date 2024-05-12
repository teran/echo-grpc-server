//go:build presenter

package presenter

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/teran/echo-grpc-server/presenter/proto"
	grpctest "github.com/teran/go-grpctest"
)

func (s *echoHandlersTestSuite) TestEcho() {
	resp, err := s.client.Echo(s.ctx, &proto.EchoRequest{
		Message: "testmessage",
	})
	s.Require().NoError(err)
	s.Require().Equal("testmessage", resp.GetMessage())
}

// ========================================================================
// Test suite setup
// ========================================================================
type echoHandlersTestSuite struct {
	suite.Suite

	srv    grpctest.Server
	ctx    context.Context
	cancel context.CancelFunc

	client   proto.EchoServiceClient
	handlers EchoHandlers
}

func (s *echoHandlersTestSuite) SetupTest() {
	s.handlers = NewEchoHandlers()

	s.srv = grpctest.New()
	s.handlers.Register(s.srv.Server())

	err := s.srv.Run()
	s.Require().NoError(err)

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 10*time.Second)

	dial, err := s.srv.DialContext(s.ctx)
	s.Require().NoError(err)

	s.client = proto.NewEchoServiceClient(dial)
}

func (s *echoHandlersTestSuite) TearDownTest() {
	s.srv.Close()
	s.cancel()
}

func TestEchoHandlersTestSuite(t *testing.T) {
	suite.Run(t, &echoHandlersTestSuite{})
}
