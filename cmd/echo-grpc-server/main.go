package main

import (
	"net"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/teran/echo-grpc-server/presenter"
)

func main() {
	var cfg Config
	envconfig.MustProcess("", &cfg)

	log.SetLevel(cfg.LogLevel)

	log.WithFields(log.Fields{
		"addr":      cfg.Addr,
		"log_level": cfg.LogLevel.String(),
	}).Debugf("running GRPC echo server")

	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		panic(err)
	}

	gs := grpc.NewServer()

	eh := presenter.NewEchoHandlers()
	eh.Register(gs)

	reh := presenter.NewRemoteEchoHandlers()
	reh.Register(gs)

	err = gs.Serve(listener)
	if err != nil {
		panic(err)
	}
}
