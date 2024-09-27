package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"tools/cmd/commands"
	"tools/pkg/logger"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "App is a command line tool for file operations",
}

func init() {
	// 初始化日志器
	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	defer logger.Cleanup() // 确保在程序退出时刷新日志

}

func main() {
	// 添加命令
	rootCmd.AddCommand(commands.RenameCmd)
	rootCmd.AddCommand(commands.RemoveAudioCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
