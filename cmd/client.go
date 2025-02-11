package cmd

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/handler"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var clientCmd = &cobra.Command{
	Use:   "query",
	Short: "Query a server at a particular endpoint",
	Long:  `Query a server at a particular endpoin`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return NewGrpcClient("127.0.0.1:8080")
	},
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

func init() {
	rootCmd.AddCommand(clientCmd)
}
