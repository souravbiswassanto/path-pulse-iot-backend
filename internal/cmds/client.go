package cmds

import (
	"context"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/cmds/server"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/handler"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewCmdClient(ctx context.Context) *cobra.Command {
	o := server.NewOptions()
	cmd := &cobra.Command{
		Use:   "client",
		Short: "Starts the grpc server",
		Long:  `Starts the grpc server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run(ctx)
		},
	}
	o.AddFlags(cmd.Flags())
	return cmd
}

func NewGrpcClient(addr string) error {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	client := handler.NewUserManagerClientHandler(conn)
	user, err := client.GetUser(uint64(1))
	if err != nil {
		return err
	}
	log.Println(*user)
	return nil
}
