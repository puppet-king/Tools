package commands

import (
	"tools/internal/custom_print"
	"tools/internal/fileutils"
	"tools/pkg/logger"

	"github.com/spf13/cobra"
)

// RemoveAudioCmd 移除音频 go run .\cmd\main.go remove-audio -s C:\BaiduSyncdisk\素材\保洁向-new
var RemoveAudioCmd = &cobra.Command{
	Use:   "remove-audio",
	Short: "remove files audio",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取命令行参数
		sourceDir, _ := cmd.Flags().GetString("source")

		err := fileutils.RemoveAudio(sourceDir)
		if err != nil {
			logger.Sugar.Errorf("Error remove files audio: %v", err)
		} else {
			custom_print.Success("Audio removed successfully from the video.")
		}
	},
}

func init() {
	// 定义命令行参数
	RemoveAudioCmd.Flags().StringP("source", "s", "C:\\BaiduSyncdisk\\素材\\美食向", "请输入来源目录 (required)")
	err := RemoveAudioCmd.MarkFlagRequired("source")
	if err != nil {
		return
	}
}
