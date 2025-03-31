package cmds

import (
	"context"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/cmds/server"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

func NewCmdServer(ctx context.Context) *cobra.Command {
	o := server.NewOptions()
	cmd := &cobra.Command{
		Use:   "start-server",
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
