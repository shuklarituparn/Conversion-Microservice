package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Watermark(inputVideo string, WatermarkImage string) string {
	inputfilePath := fmt.Sprintf("../../uploads/%s", inputVideo)
	FileWithoutExt := strings.TrimSuffix(filepath.Base(inputVideo), filepath.Ext(inputVideo))
	outputFileName := fmt.Sprintf("../userfiles/watermarked_files/%s.mp4", FileWithoutExt) //This will create a file in that location

	// Command to add watermark using FFMPEG
	cmd := exec.Command("ffmpeg",
		"-i", inputfilePath,
		"-i", WatermarkImage,
		"-filter_complex", "overlay=0:10",
		"-codec:a", "copy",
		outputFileName,
	)

	// Run the command
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return outputFileName
}
