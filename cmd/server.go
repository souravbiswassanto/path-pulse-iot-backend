package cmd

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/handler"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
)

// rootCmd represents the base command when called without any subcommands
var startServerCmd = &cobra.Command{
	Use:   "start-server",
	Short: "Starts the grpc server",
	Long:  `Starts the grpc server`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return NewGrpcServer("127.0.0.1:8080")
	},
}

func NewGrpcServer(addr string) error {
	svr := grpc.NewServer()
	h, err := handler.NewUserServerHandler()
	if err != nil {
		return err
	}
	user.RegisterUserManagerServer(svr, h)
	lsnr, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	err = svr.Serve(lsnr)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(startServerCmd)
}
