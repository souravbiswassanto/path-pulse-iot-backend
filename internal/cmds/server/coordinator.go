package server

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/in_memory"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/postgres"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/handler"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/grpc"
	"net"
)

type cancelFunc func()

type Coordinator struct {
	ctx context.Context
	*Config
	log logr.Logger
}

func (c *Coordinator) SetupAndRunServer() error {
	lsnr, err := net.Listen("tcp", c.Addr)
	if err != nil {
		return err
	}

	svr := grpc.NewServer()
	err = c.SetupAndRunUserServer(svr)
	if err != nil {
		return err
	}
	err = svr.Serve(lsnr)
	if err != nil {
		return err
	}
	return nil
}

func (c *Coordinator) SetupAndRunUserServer(svr *grpc.Server) error {
	h, cancel, err := c.SetupUserServer()
	if err != nil {
		return err
	}
	defer cancel()
	user.RegisterUserManagerServer(svr, h)
	return nil
}

func (c *Coordinator) SetupUserServer() (*handler.UserServerHandlerSer, cancelFunc, error) {
	cache := c.GetUserCacheStore()
	sqlDB, cancel, err := c.GetUserSQLDB()
	if err != nil {
		return nil, nil, err
	}
	svc := service.NewUserService(sqlDB, cache)
	h := handler.NewUserServerHandler(svc)
	return h, cancel, nil
}

func (c *Coordinator) GetUserCacheStore() *in_memory.UserInMemoryStore {
	store := in_memory.NewStore[models.UserID, *models.User](c.ctx, c.GetCacheOptions())
	return in_memory.NewUserInMemoryStore(store)
}

func (c *Coordinator) GetUserSQLDB() (*postgres.UserSqlDB, cancelFunc, error) {
	db := c.Postgres
	fn := func(pg *postgres.ClientOptions) {
		if db.Host != nil {
			pg.Host = c.Postgres.Host
		}
		if db.Port != nil {
			pg.Port = db.Port
		}
		if db.SSLMode != nil {
			pg.SSLMode = db.SSLMode
		}
		if db.Database != nil {
			pg.Database = db.Database
		}
		if db.Username != nil {
			pg.Username = db.Username
		}
		if db.Password != nil {
			pg.Password = db.Password
		}
		if db.CaCert != nil {
			pg.CaCert = db.CaCert
		}
		if db.ClientCert != nil {
			pg.ClientCert = db.ClientCert
		}
		if db.ClientKey != nil {
			pg.ClientKey = db.ClientKey
		}

		// ClientOptions fields
		if db.MaxIdleConnections != 0 {
			pg.MaxIdleConnections = db.MaxIdleConnections
		}
		if db.MaxOpenConnections != 0 {
			pg.MaxOpenConnections = db.MaxOpenConnections
		}
		if db.ConnMaxLifeTime != nil {
			pg.ConnMaxLifeTime = db.ConnMaxLifeTime
		}
	}
	client, cancel, err := postgres.NewPostgresClient(c.ctx, fn).GetPostgresXormClient()
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := postgres.NewUseSqlDB(client)
	if err != nil {
		return nil, nil, err
	}
	return sqlDB, cancel, nil
}

//func NewGrpcServer(addr string) error {
//	svr := grpc.NewServer()
//	h, err := handler.NewUserServerHandler()
//	if err != nil {
//		return err
//	}
//	user.RegisterUserManagerServer(svr, h)
//	lsnr, err := net.Listen("tcp", addr)
//	if err != nil {
//		return err
//	}
//	err = svr.Serve(lsnr)
//	if err != nil {
//		return err
//	}
//	return nil
//}
