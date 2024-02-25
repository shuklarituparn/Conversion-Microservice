package ffmpeg

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func CutFile(inputfileName string, startTime string, endTime string) string {

	inputfilePath := fmt.Sprintf("../../uploads/%s", inputfileName)
	FileWithoutExt := strings.TrimSuffix(filepath.Base(inputfileName), filepath.Ext(inputfileName))
	outputFileName := fmt.Sprintf("../userfiles/cut_files/%s.mp4", FileWithoutExt) //This will create a file in that location
	cmd := exec.Command("ffmpeg", "-i", inputfilePath, "-ss", startTime, "-to", endTime, "-c", "copy", outputFileName)
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Error cutting video: %s", err)
	}

	log.Println("Video cut successfully!")
	return outputFileName

}

//
//Without hardware accelaration it is faster for some reason
