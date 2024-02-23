package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Conversion(inputFile, outputformat string) string {

	inputfilePath := fmt.Sprintf("../../uploads/%s", inputFile)
	FileWithoutExt := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
	outputFileName := fmt.Sprintf("../userfiles/converted_files/%s.%s", FileWithoutExt, outputformat) //This will create a file in that location
	cmd := exec.Command("ffmpeg", "-i", inputfilePath, "-c:v", "h264_nvenc", outputFileName)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error converting video:", err)
		return ""
	}

	fmt.Println("Video converted successfully!")
	return outputFileName //to get it to produce a message to kafka that it is converted
}
