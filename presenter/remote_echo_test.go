//go:build presenter

package presenter

import (
	"context"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"github.com/teran/echo-grpc-server/presenter/proto"
	grpctest "github.com/teran/go-grpctest"
)

const testImage = "teran/echo-grpc-server@sha256:0d15abfab876b8f2538112fbfc09ae33fb99cc928050fa07b01d00c8bbf6ef92"

func (s *remoteEchoHandlersTestSuite) TestEcho() {
	resp, err := s.client.RemoteEcho(s.ctx, &proto.RemoteEchoRequest{
		Remote:  s.url,
		Message: "testmessage",
	})
	s.Require().NoError(err)
	s.Require().Equal("testmessage", resp.GetMessage())
}

// ========================================================================
// Test suite setup
// ========================================================================
type remoteEchoHandlersTestSuite struct {
	suite.Suite

	srv    grpctest.Server
	ctx    context.Context
	cancel context.CancelFunc

	client   proto.RemoteEchoServiceClient
	handlers RemoteEchoHandlers

	container docker.Container
	url       string
}

func (s *remoteEchoHandlersTestSuite) SetupTest() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), 30*time.Second)

	// Init remote echo server
	c, err := docker.NewContainer(
		"remote-echo-server",
		testImage,
		nil,
		docker.NewEnvironment().
			StringVar("ADDR", ":5555").
			LogLevelVar("LOG_LEVEL", log.TraceLevel),
		docker.NewPortBindings().
			PortDNAT(docker.ProtoTCP, 5555),
	)
	s.Require().NoError(err)

	s.container = c

	u, err := s.container.URL(docker.ProtoTCP, 5555)
	s.Require().NoError(err)

	s.url = u.String()

	err = s.container.Run(s.ctx)
	s.Require().NoError(err)

	err = s.container.AwaitOutput(s.ctx, docker.NewSubstringMatcher("running GRPC echo server"))
	s.Require().NoError(err)

	// Init remote echo server
	s.handlers = NewRemoteEchoHandlers()

	s.srv = grpctest.New()
	s.handlers.Register(s.srv.Server())

	err = s.srv.Run()
	s.Require().NoError(err)

	dial, err := s.srv.DialContext(s.ctx)
	s.Require().NoError(err)

	// init remote echo client
	s.client = proto.NewRemoteEchoServiceClient(dial)
}

func (s *remoteEchoHandlersTestSuite) TearDownTest() {
	s.container.Close(s.ctx)
	s.srv.Close()
	s.cancel()
}

func TestRemoteEchoHandlersTestSuite(t *testing.T) {
	suite.Run(t, &remoteEchoHandlersTestSuite{})
}
