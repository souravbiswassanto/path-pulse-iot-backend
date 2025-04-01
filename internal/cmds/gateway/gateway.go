package gateway

import (
	"context"
	"github.com/souravbiswassanto/path-pulse-iot-backend/pkg/lib"
	"github.com/spf13/pflag"
)

type Options struct {
	Addr, ServerAddr, MetricsAddr string
	Verbosity                     int
}

func NewOptions() *Options {
	return &Options{
		ServerAddr:  "127.0.0.1:8978",
		MetricsAddr: "127.0.0.1:8979",
		Addr:        "127.0.0.1:8980",
		Verbosity:   3,
	}
}

func (o *Options) Run(ctx context.Context) error {
	c := Coordinator{
		ctx:    ctx,
		Config: o.GetConfig(),
		log:    lib.NewLogger(o.Verbosity).WithName("gateway"),
	}
	return c.SetupGateway()
}

func (o *Options) Validate() error {

	return nil
}

func (o *Options) Complete() error {
	return nil
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Addr, "gw-address", o.Addr, "The Grpc Gateway Address")
	fs.StringVar(&o.ServerAddr, "server-address", o.ServerAddr, "The Grpc Server Address")
	fs.StringVar(&o.MetricsAddr, "metrics-address", o.MetricsAddr, "The Metrics Server Address")
	fs.IntVar(&o.Verbosity, "v", o.Verbosity, "Number for the log level verbosity")

}
