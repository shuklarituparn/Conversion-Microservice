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

func Watermark(inputVideo string, WatermarkImage string) string {
	correctFilePath := fmt.Sprintf("../%s", inputVideo)
	FileWithoutExt := strings.TrimSuffix(filepath.Base(inputVideo), filepath.Ext(inputVideo))
	outputFileName := fmt.Sprintf("../../internal/userfiles/watermarked_files/%s.mp4", FileWithoutExt) //This will create a file in that location

	// Command to add watermark using FFMPEG
	cmd := exec.Command("ffmpeg",
		"-i", correctFilePath,
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

func WatermarkConsumer() {

	c, _ := consumer.NewConsumer("localhost:9092", "conversion_service")
	_ = c.Subscribe("watermark_video", nil)

	defer consumer.Close(c)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		var wtrmrkMsg models.WatermarkMessage
		err = json.Unmarshal(msg.Value, &wtrmrkMsg)
		if err != nil {
			fmt.Println(err)
		}
		outputFilePath := Watermark(wtrmrkMsg.FilePath, wtrmrkMsg.WaterMarkFile) //we get the file path here

		// Need to produce a message on the topic to upload
		//Prodcue the message for the mongo consumer with Video Key
		var messageToUpload models.AfterCutUpload
		outputFilName := filepath.Base(outputFilePath)
		messageToUpload.VideoKey = wtrmrkMsg.VideoKey
		messageToUpload.FilePath = outputFilePath
		messageToUpload.UserId = wtrmrkMsg.UserId
		messageToUpload.FileName = outputFilName
		messageToUpload.UserName = wtrmrkMsg.UserName
		serializedMessage, errorSerializing := json.Marshal(messageToUpload)
		if errorSerializing != nil {
			log.Println("Error serializing message", errorSerializing)

		}
		p, err := producer.NewProducer("localhost:9092")
		producer.ProduceNewMessage(p, "upload", string(serializedMessage)) //Producing the converted mesaage
		//Now need the upload consumer
		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}
