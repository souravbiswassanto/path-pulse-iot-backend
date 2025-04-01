package server

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/in_memory"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/influx"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/postgres"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	*DatabaseConfig
	Addr, MetricsAddr  string
	DatabaseConfigPath string
	Verbosity          int
}

type DatabaseConfig struct {
	Postgres *postgres.ClientOptions `yaml:"postgres,omitempty"`
	Influx   *influx.InfluxDBOptions `yaml:"influx,omitempty"`
	Cache    *in_memory.StoreOptions `yaml:"cache,omitempty"`
}

func (o *Options) GetConfig() (*Config, error) {
	cfg := &Config{}
	cfg.Verbosity = o.Verbosity
	cfg.Addr = o.Addr
	cfg.MetricsAddr = o.MetricsAddr
	cfg.DatabaseConfigPath = o.DatabaseConfigPath
	dbc, err := cfg.ParseDatabaseConfig()
	if err != nil {
		return nil, err
	}
	cfg.DatabaseConfig = dbc
	return cfg, nil
}

func (c *Config) ParseDatabaseConfig() (*DatabaseConfig, error) {
	if c.DatabaseConfigPath == "" {
		return &DatabaseConfig{}, nil
	}
	out, err := os.ReadFile(c.DatabaseConfigPath)
	if err != nil {
		return &DatabaseConfig{}, err
	}
	dc := &DatabaseConfig{}
	err = yaml.Unmarshal(out, dc)
	if err != nil {
		return &DatabaseConfig{}, err
	}
	return dc, nil
}

func (dc *DatabaseConfig) GetCacheOptions() *in_memory.StoreOptions {
	if dc.Cache == nil || dc.Cache.MaxKeyStoreLimit == 0 {
		return &in_memory.StoreOptions{
			MaxKeyStoreLimit: 100000,
		}
	}
	return dc.Cache
}

func (dc *DatabaseConfig) GetPostgresOptions() *postgres.ClientOptions {
	if dc.Postgres == nil {
		return &postgres.ClientOptions{}
	}
	return dc.Postgres
}
func (dc *DatabaseConfig) GetInfluxDBOptions() *influx.InfluxDBOptions {
	if dc.Influx == nil {
		return &influx.InfluxDBOptions{}
	}
	return dc.Influx
}
