package postgres

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	passgen "gomodules.xyz/password-generator"
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
	maxIdleConnections int
	maxOpenConnections int
	connMaxLifeTime    *time.Duration
}

func NewPostgresClient(opts ...func(*ClientOptions)) *PostgresClient {
	co := &ClientOptions{
		Connector: &Connector{},
		ctx:       context.TODO(),
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
	co.maxIdleConnections = 10
	co.maxOpenConnections = 15
	t := time.Minute * 30
	co.connMaxLifeTime = &t
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
	engine.SetMaxIdleConns(c.maxIdleConnections)
	engine.SetMaxOpenConns(c.maxOpenConnections)
	if c.connMaxLifeTime != nil {
		engine.SetConnMaxLifetime(*c.connMaxLifeTime)
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
	ctx        context.Context
	port       *int
	host       *string
	sslMode    *string
	database   *string
	username   *string
	password   *string
	caCert     *string
	clientCert *string
	clientKey  *string
}

func (c *Connector) WithPort(port int) *Connector {
	*c.port = port
	return c
}
func (c *Connector) WithHost(host string) *Connector {
	*c.host = host
	return c
}
func (c *Connector) WithSSLMode(sslMode string) *Connector {
	*c.sslMode = sslMode
	return c
}
func (c *Connector) WithDatabase(database string) *Connector {
	*c.database = database
	return c
}
func (c *Connector) WithUserName(username string) *Connector {
	*c.username = username
	return c
}
func (c *Connector) WithPassword(password string) *Connector {
	*c.password = password
	return c
}
func (c *Connector) WithCaCert(cert string) *Connector {
	*c.caCert = cert
	return c
}
func (c *Connector) WithClientCert(cert string) *Connector {
	*c.clientCert = cert
	return c
}
func (c *Connector) WithClientKey(key string) *Connector {
	*c.clientKey = key
	return c
}
func (c *Connector) WithContext(ctx context.Context) *Connector {
	c.ctx = ctx
	return c
}
func (c *Connector) BuildConnectionString() string {
	connstr := ""
	username := "postgres"
	if c.username != nil {
		username = *c.username
	}
	connstr += "user=" + username + " "
	if c.password != nil {
		connstr += "password=" + *c.password + " "
	} else {
		connstr += "password=" + passgen.Generate(8) + " "
	}
	host := "127.0.0.1"
	if c.host != nil {
		host = *c.host
	}
	connstr += "host=" + host + " "
	port := 5432
	if c.port != nil {
		port = *c.port
	}
	connstr += "port=" + fmt.Sprintf("%d", port) + " "
	database := "postgres"
	if c.database != nil {
		database = *c.database
	}
	connstr += "dbname=" + database + " "
	sslMode := "disable"
	if c.sslMode != nil {
		sslMode = *c.sslMode
	}
	connstr += "sslmode=" + sslMode + " "
	if c.caCert != nil {
		connstr += "sslrootcert==" + *c.caCert + " "
	}
	if c.clientCert != nil {
		connstr += "sslcert==" + *c.clientCert + " "
	}
	if c.clientKey != nil {
		connstr += "sslkey==" + *c.clientKey + " "
	}
	return connstr
}
