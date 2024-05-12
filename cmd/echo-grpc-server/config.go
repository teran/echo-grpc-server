package main

import (
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Addr     string    `envconfig:"addr" required:"true"`
	LogLevel log.Level `envconfig:"log_level" default:"info"`
}
