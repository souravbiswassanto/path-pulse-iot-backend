package server

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/in_memory"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/postgres"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/handler"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/event"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/group"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
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
	err = c.SetupServer(svr)
	if err != nil {
		return err
	}
	err = svr.Serve(lsnr)
	if err != nil {
		return err
	}
	return nil
}

func (c *Coordinator) GetSQLDB(v interface{}) (interface{}, cancelFunc, error) {
	client, cancel, err := postgres.NewPostgresClient(c.ctx, getPostgresClientOptions(c.Postgres)).GetPostgresXormClient()
	if err != nil {
		return nil, nil, err
	}

	var sqlDB interface{}
	switch v.(type) {
	case *postgres.UserSqlDB:
		sqlDB, err = postgres.NewUserSqlDB(client)
	case *postgres.EventSqlDB:
		sqlDB, err = postgres.NewEventSqlDB(client)
	case *postgres.GroupSqlDB:
		sqlDB, err = postgres.NewGroupSqlDB(client)
	default:
		err = fmt.Errorf("unkown type. expected any of %T, %T, %T, got %T", &postgres.UserSqlDB{}, &postgres.EventSqlDB{}, &postgres.GroupSqlDB{}, v)
	}
	if err != nil {
		return nil, nil, err
	}
	return sqlDB, cancel, nil
}

func (c *Coordinator) SetupServer(svr *grpc.Server) error {
	err := c.SetupAndRegisterUserServer(svr)
	if err != nil {
		return err
	}
	err = c.SetupAndRegisterEventServer(svr)
	if err != nil {
		return err
	}
	err = c.SetupAndRegisterGroupServer(svr)
	if err != nil {
		return err
	}
	err = c.SetupAndRegisterTrackerServer(svr)
	if err != nil {
		return err
	}
	return nil
}

func (c *Coordinator) SetupAndRegisterUserServer(svr *grpc.Server) error {
	h, err := c.SetupUserServer()
	if err != nil {
		return err
	}
	user.RegisterUserManagerServer(svr, h)
	return nil
}

func (c *Coordinator) SetupUserServer() (*handler.UserServerHandlerSer, error) {
	cache := c.GetUserCacheStore()
	sqlDB, cancel, err := c.GetSQLDB(&postgres.UserSqlDB{})
	if err != nil {
		return nil, err
	}
	go func() {
		select {
		case <-c.ctx.Done():
			cancel()
		}
	}()
	svc := service.NewUserService(sqlDB.(*postgres.UserSqlDB), cache)
	h := handler.NewUserServerHandler(svc)
	return h, nil
}

func (c *Coordinator) GetUserCacheStore() *in_memory.UserInMemoryStore {
	store := in_memory.NewStore[models.UserID, *models.User](c.ctx, c.GetCacheOptions())
	return in_memory.NewUserInMemoryStore(store)
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

func (c *Coordinator) SetupAndRegisterEventServer(svr *grpc.Server) error {
	h, err := c.SetupEventServer()
	if err != nil {
		return err
	}
	event.RegisterEventManagerServer(svr, h)
	return nil
}

func (c *Coordinator) SetupEventServer() (*handler.EventServerHandler, error) {
	cache := c.GetEventCacheStore()
	sqlDB, cancel, err := c.GetSQLDB(&postgres.EventSqlDB{})
	if err != nil {
		return nil, err
	}
	go func() {
		select {
		case <-c.ctx.Done():
			cancel()
		}
	}()
	svc := service.NewEventService(sqlDB.(*postgres.EventSqlDB), cache)
	h := handler.NewEventServerHandler(svc)
	return h, nil
}

func (c *Coordinator) GetEventCacheStore() *in_memory.EventInMemoryStore {
	store := in_memory.NewStore[uint64, *models.Event](c.ctx, c.GetCacheOptions())
	return in_memory.NewEventInMemoryStore(store)
}

func (c *Coordinator) SetupAndRegisterGroupServer(svr *grpc.Server) error {
	h, err := c.SetupGroupServer()
	if err != nil {
		return err
	}
	group.RegisterGroupManagerServer(svr, h)
	return nil
}

func (c *Coordinator) SetupGroupServer() (*handler.GroupServerHandler, error) {
	cache := c.GetGroupCacheStore()
	sqlDB, cancel, err := c.GetSQLDB(&postgres.GroupSqlDB{})
	if err != nil {
		return nil, err
	}
	go func() {
		select {
		case <-c.ctx.Done():
			cancel()
		}
	}()
	svc := service.NewGroupService(sqlDB.(*postgres.GroupSqlDB), cache)
	h := handler.NewGroupServerHandler(svc)
	return h, nil
}

func (c *Coordinator) GetGroupCacheStore() *in_memory.GroupInMemoryStore {
	store := in_memory.NewStore[uint64, *models.Group](c.ctx, c.GetCacheOptions())
	return in_memory.NewGroupInMemoryStore(store)
}

func (c *Coordinator) SetupAndRegisterTrackerServer(svr *grpc.Server) error {
	h, err := c.SetupTrackerServer()
	if err != nil {
		return err
	}
	tracker.RegisterTrackerServer(svr, h)
	return nil
}

func (c *Coordinator) SetupTrackerServer() (*handler.TrackerHandlerServer, error) {
	svc := service.NewTrackerService(c.GetInfluxDBOptions())
	h := handler.NewTrackerServerHandler(svc)
	return h, nil
}

func getPostgresClientOptions(db *postgres.ClientOptions) func(options *postgres.ClientOptions) {
	return func(pg *postgres.ClientOptions) {
		if db.Host != nil {
			pg.Host = db.Host
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
}
