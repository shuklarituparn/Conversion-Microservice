package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Screenshot(videoFileName, timeScreeshot string) string {
	inputFIlepath := fmt.Sprintf("../../uploads/%s", videoFileName)
	FileWithoutExt := strings.TrimSuffix(filepath.Base(videoFileName), filepath.Ext(videoFileName))
	outputScreenshot := fmt.Sprintf("../userfiles/screenshot_files/%s.jpg", FileWithoutExt)
	cmd := exec.Command("ffmpeg", "-ss", timeScreeshot, "-i", inputFIlepath, "-frames:v", "1", "-q:v", "2", outputScreenshot)

	// Execute the command
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	fmt.Println("Screenshot taken successfully!")

	return outputScreenshot
}
