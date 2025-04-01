/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/cmds"
	"gomodules.xyz/logs"
	"k8s.io/klog/v2"
)

func main() {
	rootCmd := cmds.NewRootCmd()
	logs.Init(rootCmd, true)
	defer logs.FlushLogs()

	if err := rootCmd.Execute(); err != nil {
		klog.Warningln(err)
	}
}
