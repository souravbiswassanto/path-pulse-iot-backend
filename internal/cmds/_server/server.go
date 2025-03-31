package server

import (
	"context"
	"github.com/souravbiswassanto/path-pulse-iot-backend/pkg/lib"
	"github.com/spf13/pflag"
)

const (
	defaultDatabaseConfigPath = "internal/config/database-config.yaml"
)

type Options struct {
	Addr, MetricsAddr  string
	DatabaseConfigPath string
	Verbosity          int
}

func NewOptions() *Options {
	return &Options{
		Addr:               "127.0.0.1:8978",
		MetricsAddr:        "127.0.0.1:8979",
		DatabaseConfigPath: defaultDatabaseConfigPath,
		Verbosity:          3,
	}
}

func (o *Options) Run(ctx context.Context) error {
	config, err := o.GetConfig()
	if err != nil {
		return err
	}
	c := Coordinator{
		ctx:    ctx,
		Config: config,
		log:    lib.NewLogger(o.Verbosity).WithName("server"),
	}
	return c.SetupAndRunServer()
}

func (o *Options) Validate() error {

	return nil
}

func (o *Options) Complete() error {
	return nil
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	//fs.Float64Var(&s.QPS, "qps", s.QPS, "The maximum QPS to the master from this client")
	//fs.IntVar(&s.Burst, "burst", s.Burst, "The maximum burst for throttle")
	//
	//fs.BoolVar(&s.EnableMutatingWebhook, "enable-mutating-webhook", s.EnableMutatingWebhook, "If true, enables mutating webhooks for KubeDB CRDs.")
	//fs.BoolVar(&s.EnableValidatingWebhook, "enable-validating-webhook", s.EnableValidatingWebhook, "If true, enables validating webhooks for KubeDB CRDs.")
	//fs.StringVar(&s.MetricsAddr, "metrics-bind-address", s.MetricsAddr, "The address the metric endpoint binds to. "+
	//	"Use the port :8080. If not set, it will be 0 in order to disable the metrics server")
	//fs.BoolVar(&s.SecureMetrics, "metrics-secure", s.SecureMetrics,
	//	"If set the metrics endpoint is served securely")
	//fs.BoolVar(&s.EnableHTTP2, "enable-http2", s.EnableHTTP2,
	//	"If set, HTTP/2 will be enabled for the metrics and webhook servers")

	fs.StringVar(&o.Addr, "server-address", o.Addr, "The Grpc Server Address")
	fs.StringVar(&o.MetricsAddr, "metrics-address", o.MetricsAddr, "The Metrics Server Address")
	fs.StringVar(&o.DatabaseConfigPath, "db-config-path", o.DatabaseConfigPath, "Database Configuration File Path")
	fs.IntVar(&o.Verbosity, "v", o.Verbosity, "Number for the log level verbosity")

}
