package fileutils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"tools/internal/custom_print"
	"tools/pkg/logger"
)

// RenameFilesBasedOnMetadata 根据文件的属性重命名文件
func RenameFilesBasedOnMetadata(sourceDir, targetDir string) error {
	counter := 0

	// 如果 targetDir 为空，则使用 sourceDir 作为目标目录
	if targetDir == "" {
		targetDir = sourceDir
	}

	targetDir = createDir(targetDir)
	videoDir := createDir(targetDir + "/video")
	imgDir := createDir(targetDir + "/img")

	cur, err := os.ReadDir(sourceDir)
	if err != nil {
		logger.Sugar.Errorf("读取目录时出错：%v", err)
		return err
	}

	// 遍历源目录中的所有文件
	for _, filename := range cur {
		// 获取文件的修改时间（此处示例使用修改时间作为创建时间）
		if filename.IsDir() {
			break
		}

		info, _ := filename.Info()
		creationTime := info.ModTime()

		// 根据创建日期格式化新文件名
		dateStr := creationTime.Format("0102")
		newName := fmt.Sprintf("%s-%d%s", dateStr, counter, filepath.Ext(info.Name()))

		// 生成目标路径
		newPath := filepath.Join(targetDir, newName)
		if isVideo(info.Name()) {
			newPath = filepath.Join(videoDir, newName)
		} else if isImg(info.Name()) {
			newPath = filepath.Join(imgDir, newName)
		}

		// 移动并重命名文件
		err = os.Rename(filepath.Join(sourceDir, filename.Name()), newPath)
		if err != nil {
			custom_print.Error("failed to rename and move file: %v", err)
		}

		custom_print.Info("Renamed and moved %s to %s", filepath.Join(sourceDir, filename.Name()), newPath)
		counter++
	}

	return nil
}

// RemoveAudio 从视频文件中移除音频
func RemoveAudio(sourceDir string) error {
	// 确保目标目录存在，如果不存在则创建
	outputDir := createDir(sourceDir + "/remove-audio")

	cur, err := os.ReadDir(sourceDir)
	if err != nil {
		logger.Sugar.Errorf("读取目录时出错：%v", err)
		return err
	}

	// 遍历源目录中的所有文件
	for _, info := range cur {
		// 忽略子目录，只处理文件
		if info.IsDir() {
			break
		}

		inputFile := filepath.Join(sourceDir, info.Name())
		outputFile := filepath.Join(outputDir, info.Name())

		if isVideo(info.Name()) {
			// 构建 FFmpeg 命令：-an 表示去除音频
			cmd := exec.Command(getFfmpeg(), "-i", inputFile, "-an", outputFile, "-y")
			_, err := cmd.CombinedOutput()
			if err != nil {
				custom_print.Error("文件 %s 转码失败", inputFile)
				return fmt.Errorf("failed to remove audio: %v", err)
			} else {
				custom_print.Info("文件 %s 转码成功", inputFile)
			}
		}

	}

	return nil
}

// getFfmpeg 获取项目的 Ffmpeg 路径
func getFfmpeg() string {
	// 获取当前工作目录（即程序运行时的路径）
	currentDir, _ := os.Getwd()

	ffmpegPath := filepath.Join(currentDir, "ffmpeg")

	if FileExists(ffmpegPath + ".exe") {
		return ffmpegPath
	} else {
		custom_print.Error("找不到 Ffmpeg, 请确认在根目录下")
		os.Exit(-1)
		return ""
	}
}

// FileExists 判断文件或目录是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		// 如果没有错误，表示文件存在
		return true
	}
	if os.IsNotExist(err) {
		// 如果错误类型是文件不存在
		return false
	}

	// 其他错误，比如权限问题
	logger.Sugar.Errorf("Error checking file: %v\n", err)
	return false
}

// isImg 是否是图片
func isImg(file string) bool {
	imgExtensions := []string{
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".tif",
		".webp", ".svg", ".heic", ".heif", ".ico", ".raw", ".psd",
		".eps", ".ai", ".tga", ".jp2",
	}

	// 获取文件的扩展名并转换为小写
	ext := strings.ToLower(filepath.Ext(file))

	// 检查文件扩展名是否在视频扩展名列表中
	for _, imgExt := range imgExtensions {
		if ext == imgExt {
			return true
		}
	}

	return false
}

// isVideo 是否是视频
func isVideo(file string) bool {
	videoExtensions := []string{".mp4", ".mov", ".mkv", ".avi", ".flv", ".wmv", ".webm", ".mpg", ".mpeg"}

	// 获取文件的扩展名并转换为小写
	ext := strings.ToLower(filepath.Ext(file))

	// 检查文件扩展名是否在视频扩展名列表中
	for _, videoExt := range videoExtensions {
		if ext == videoExt {
			return true
		}
	}

	return false
}

// createDir 创建目录
func createDir(dir string) string {
	_, err := os.Stat(dir)
	if err == nil {
		// 如果没有错误，表示文件存在
		return dir
	}

	if os.IsNotExist(err) {
		// 确保目标目录存在，如果不存在则创建
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			logger.Sugar.Errorf("failed to create target directory: %v", err)
			return ""
		}

		custom_print.Trace("创建目录成功  路径为：%s", dir)
		return dir
	} else {
		logger.Sugar.Errorf("failed to create target directory: %v", err)
		return ""
	}
}
