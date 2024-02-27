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

func CutFile(inputfileName string, startTime string, endTime string) string {

	inputfilePath := fmt.Sprintf("../../uploads/%s", inputfileName)
	FileWithoutExt := strings.TrimSuffix(filepath.Base(inputfileName), filepath.Ext(inputfileName))
	outputFileName := fmt.Sprintf("../../internal/userfiles/cut_files/%s.mp4", FileWithoutExt) //This will create a file in that location
	cmd := exec.Command("ffmpeg", "-y", "-i", inputfilePath, "-ss", startTime, "-to", endTime, "-c", "copy", outputFileName)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error cutting video: %s", err)
	}

	log.Println("Video cut successfully!")
	return outputFileName

}

func VideoCutConsumer() {

	c, _ := consumer.NewConsumer("localhost:9092", "conversion_service")
	_ = c.Subscribe("cut_video", nil)

	defer consumer.Close(c)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		var cutMsg models.CutMessage
		err = json.Unmarshal(msg.Value, &cutMsg)
		if err != nil {
			fmt.Println(err)
		}

		outputFilePath := CutFile(cutMsg.FileName, cutMsg.StartTime, cutMsg.EndTime) //we get the file path here

		// Need to produce a message on the topic to upload
		//Prodcue the message for the mongo consumer with Video Key
		var messageToUpload models.AfterCutUpload
		FileName := filepath.Base(outputFilePath)
		messageToUpload.VideoKey = cutMsg.VideoKey
		messageToUpload.FilePath = outputFilePath
		messageToUpload.UserId = cutMsg.UserId
		messageToUpload.FileName = FileName
		messageToUpload.UserName = cutMsg.UserName
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
