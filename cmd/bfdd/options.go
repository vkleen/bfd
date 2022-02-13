package main

import (
	"github.com/spf13/pflag"
)

type Options struct {
	HelpRequested bool
	Config        string
  GrpcSocket    string
}

func NewOptions() *Options {
	return &Options{
		Config: "/etc/bfdd/config.yaml",
	}
}

func (s *Options) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVarP(&s.HelpRequested, "help", "h", false, "Print usage information.")
	fs.StringVarP(&s.Config, "config", "c", s.Config, "The path to the configuration file.")
  fs.StringVarP(&s.GrpcSocket, "grpc-socket", "s", "/tmp/bfdd.sock", "The path to the gRPC unix socket")
}
