/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmds

import (
	"context"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

func NewRootCmd() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "path-pulse-iot-backend",
		Short: "Path Pulse IOT Backend",
		Long:  `Path Pulse IOT Backend`,
	}
	// ensuring termination of all go routines
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	rootCmd.AddCommand(NewCmdServer(ctx))
	rootCmd.AddCommand(NewCmdGateway(ctx))
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return rootCmd
}
