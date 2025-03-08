package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	*sql.DB
}

func NewPostgresClient(connstr string) (*PostgresClient, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}
	return &PostgresClient{db}, nil
}

type Connector struct {
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

func NewPostgresConnector() *Connector {
	return &Connector{}
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
func (c *Connector) BuildConnectionString() string {
	connstr := ""
	username := "postgres"
	if c.username != nil {
		username = *c.username
	}
	connstr += "user=" + username + " "
	if c.password != nil {
		connstr += "password=" + *c.password + " "
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
