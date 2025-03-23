package postgres

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	passgen "gomodules.xyz/password-generator"
	"os"
	"strconv"
	"time"
	"xorm.io/xorm"
)

type PostgresClient struct {
	*ClientOptions
	db *sql.DB
	xc *xorm.Engine
}

type ClientOptions struct {
	ctx context.Context
	*Connector
	MaxIdleConnections int            `yaml:"maxIdleConnections,omitempty"`
	MaxOpenConnections int            `yaml:"maxOpenConnections,omitempty"`
	ConnMaxLifeTime    *time.Duration `yaml:"connMaxLifeTime,omitempty"`
}

func NewPostgresClient(ctx context.Context, opts ...func(*ClientOptions)) *PostgresClient {
	co := &ClientOptions{
		Connector: &Connector{},
		ctx:       ctx,
	}
	co.SetDefaultConnectionPooling()
	for _, f := range opts {
		f(co)
	}
	return &PostgresClient{
		ClientOptions: co,
	}
}

func (co *ClientOptions) SetDefaultConnectionPooling() {
	co.MaxIdleConnections = 10
	co.MaxOpenConnections = 15
	t := time.Minute * 30
	co.ConnMaxLifeTime = &t
}

// GetPostgresXormClient Creates a xorm client, returns a cancel func
func (c *PostgresClient) GetPostgresXormClient() (*PostgresClient, func(), error) {
	if c.ctx == nil {
		c.ctx = context.Background()
	}
	connector := c.BuildConnectionString()

	engine, err := xorm.NewEngine("postgres", connector)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate postgres client using connection string: %v", err)
	}
	_, err = engine.Query("SELECT 1")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run query: %v", err)
	}
	engine.SetDefaultContext(c.ctx)
	// https://xorm.io/docs/chapter-01/readme/#connections-pool
	engine.SetMaxIdleConns(c.MaxIdleConnections)
	engine.SetMaxOpenConns(c.MaxOpenConnections)
	if c.ConnMaxLifeTime != nil {
		engine.SetConnMaxLifetime(*c.ConnMaxLifeTime)
	}
	return &PostgresClient{xc: engine}, func() {
		_ = engine.Close()
	}, nil
}

func (c *Connector) GetPostgresDBClient() (*PostgresClient, error) {
	db, err := sql.Open("postgres", c.BuildConnectionString())
	if err != nil {
		return nil, err
	}
	return &PostgresClient{db: db}, nil
}

type Connector struct {
	Port       *int    `yaml:"port,omitempty"`
	Host       *string `yaml:"host,omitempty"`
	SSLMode    *string `yaml:"sslmode,omitempty"`
	Database   *string `yaml:"database,omitempty"`
	Username   *string `yaml:"username,omitempty"`
	Password   *string `yaml:"password,omitempty"`
	CaCert     *string `yaml:"caCert,omitempty"`
	ClientCert *string `yaml:"clientCert,omitempty"`
	ClientKey  *string `yaml:"clientKey,omitempty"`
}

func (c *Connector) WithPort(port int) *Connector {
	*c.Port = port
	return c
}
func (c *Connector) WithHost(Host string) *Connector {
	*c.Host = Host
	return c
}
func (c *Connector) WithSSLMode(SSLMode string) *Connector {
	*c.SSLMode = SSLMode
	return c
}
func (c *Connector) WithDatabase(Database string) *Connector {
	*c.Database = Database
	return c
}
func (c *Connector) WithUserName(Username string) *Connector {
	*c.Username = Username
	return c
}
func (c *Connector) WithPassword(Password string) *Connector {
	*c.Password = Password
	return c
}
func (c *Connector) WithCaCert(cert string) *Connector {
	*c.CaCert = cert
	return c
}
func (c *Connector) WithClientCert(cert string) *Connector {
	*c.ClientCert = cert
	return c
}
func (c *Connector) WithClientKey(key string) *Connector {
	*c.ClientKey = key
	return c
}

func (c *Connector) BuildConnectionString() string {
	connstr := ""
	Username := "postgres"
	if c.Username != nil {
		Username = *c.Username
	} else if os.Getenv("PGUSER") != "" {
		Username = os.Getenv("PGUSER")
	}
	connstr += "user=" + Username + " "
	Password := ""
	if c.Password != nil {
		Password = *c.Password
	} else if os.Getenv("PGPASSWORD") != "" {
		Password = os.Getenv("PGPASSWORD")
	} else {
		Password = passgen.Generate(8) // Fallback to generated Password
	}
	connstr += "Password=" + Password + " "

	Host := "127.0.0.1"
	if c.Host != nil {
		Host = *c.Host
	} else if os.Getenv("PGHOST") != "" {
		Host = os.Getenv("PGHOST")
	}
	connstr += "Host=" + Host + " "

	Port := 5432
	if c.Port != nil {
		Port = *c.Port
	} else if os.Getenv("PGPORT") != "" {
		Port, _ = strconv.Atoi(os.Getenv("PGPORT")) // Convert string to int
	}
	connstr += "Port=" + fmt.Sprintf("%d", Port) + " "

	Database := "postgres"
	if c.Database != nil {
		Database = *c.Database
	} else if os.Getenv("PGDATABASE") != "" {
		Database = os.Getenv("PGDATABASE")
	}
	connstr += "dbname=" + Database + " "

	SSLMode := "disable"
	if c.SSLMode != nil {
		SSLMode = *c.SSLMode
	} else if os.Getenv("PGSSLMODE") != "" {
		SSLMode = os.Getenv("PGSSLMODE")
	}
	connstr += "sslmode=" + SSLMode + " "

	if c.CaCert != nil {
		connstr += "sslrootcert=" + *c.CaCert + " "
	} else if os.Getenv("PGSSLROOTCERT") != "" {
		connstr += "sslrootcert=" + os.Getenv("PGSSLROOTCERT") + " "
	}

	if c.ClientCert != nil {
		connstr += "sslcert=" + *c.ClientCert + " "
	} else if os.Getenv("PGSSLCERT") != "" {
		connstr += "sslcert=" + os.Getenv("PGSSLCERT") + " "
	}

	if c.ClientKey != nil {
		connstr += "sslkey=" + *c.ClientKey + " "
	} else if os.Getenv("PGSSLKEY") != "" {
		connstr += "sslkey=" + os.Getenv("PGSSLKEY") + " "
	}
	return connstr
}
