package ffmpeg

import (
	"encoding/json"
	"fmt"
	"github.com/shuklarituparn/Conversion-Microservice/internal/consumer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Screenshot(videoFileName, timeScreeshot string) string {
	inputFIlepath := fmt.Sprintf("../../uploads/%s", videoFileName)
	FileWithoutExt := strings.TrimSuffix(filepath.Base(videoFileName), filepath.Ext(videoFileName))
	outputScreenshot := fmt.Sprintf("../../internal/userfiles/screenshot_files/%s.jpg", FileWithoutExt)
	cmd := exec.Command("ffmpeg", "-y", "-ss", timeScreeshot, "-i", inputFIlepath, "-frames:v", "1", "-q:v", "2", outputScreenshot)

	// Execute the command
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	fmt.Println("Screenshot taken successfully!")

	return outputScreenshot
}

func ScreenshotConsumer() {

	c, _ := consumer.NewConsumer("localhost:9092", "conversion_service")
	_ = c.Subscribe("screenshot_video", nil)

	defer consumer.Close(c)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		var scrnShotMsg models.ScreenshotMessage
		err = json.Unmarshal(msg.Value, &scrnShotMsg)
		if err != nil {
			fmt.Println(err)
		}
		outputFilePath := Screenshot(scrnShotMsg.FileName, scrnShotMsg.Time) //we get the file path here

		// Need to produce a message on the topic to upload
		//Prodcue the message for the mongo consumer with Video Key
		var messageToUpload models.AfterScreenshotUpload
		FileName := filepath.Base(outputFilePath)
		messageToUpload.VideoKey = scrnShotMsg.VideoKey
		messageToUpload.FilePath = outputFilePath
		messageToUpload.UserId = scrnShotMsg.UserId
		messageToUpload.FileName = FileName
		messageToUpload.UserName = scrnShotMsg.UserName
		serializedMessage, errorSerializing := json.Marshal(messageToUpload)
		if errorSerializing != nil {
			log.Println("Error serializing message", errorSerializing)

		}
		p, err := producer.NewProducer("localhost:9092")
		producer.ProduceNewMessage(p, "upload_screenshot", string(serializedMessage)) //Producing the converted mesaage
		//Now need the upload consumer
		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}
