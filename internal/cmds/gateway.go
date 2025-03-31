package cmds

import (
	"context"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/cmds/gateway"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

func NewCmdGateway(ctx context.Context) *cobra.Command {
	o := gateway.NewOptions()
	cmd := &cobra.Command{
		Use:   "start-gw",
		Short: "Starts the grpc gateway client",
		Long:  `Starts the grpc gateway client`,
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
