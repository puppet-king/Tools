package commands

import (
	"tools/internal/custom_print"
	"tools/internal/fileutils"
	"tools/pkg/logger"

	"github.com/spf13/cobra"
)

// RenameCmd 定义子命令 go run .\cmd\main.go rename -s C:\BaiduSyncdisk\素材\保洁向
var RenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files based on metadata",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取命令行参数
		sourceDir, _ := cmd.Flags().GetString("source")
		targetDir, _ := cmd.Flags().GetString("target")

		// 调用文件重命名函数
		err := fileutils.RenameFilesBasedOnMetadata(sourceDir, targetDir)
		if err != nil {
			custom_print.Error("%v", err)
			logger.Sugar.Errorf("Error renaming files: %v", err)
		} else {
			custom_print.Success("All files have been renamed and moved successfully!")
		}
	},
}

func init() {
	// 定义命令行参数
	RenameCmd.Flags().StringP("source", "s", "C:\\BaiduSyncdisk\\素材\\美食向", "请输入来源目录 (required)")
	RenameCmd.Flags().StringP("target", "t", "", "请输入输出目录")
	//err := RenameCmd.MarkFlagRequired("source")
	//if err != nil {
	//	return
	//}
}
