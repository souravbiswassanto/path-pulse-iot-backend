package gateway

type Config struct {
	GrpcAddr, Addr, MetricsAddr string
	Verbosity                   int
}

func (o *Options) GetConfig() *Config {
	cfg := &Config{}
	cfg.Verbosity = o.Verbosity
	cfg.Addr = o.Addr
	cfg.MetricsAddr = o.MetricsAddr
	return cfg
}
