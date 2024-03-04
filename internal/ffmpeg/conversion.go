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

func Conversion(inputFile, outputformat string) string {
	inputfilePath := fmt.Sprintf("../../uploads/%s", inputFile)
	FileWithoutExt := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
	outputFileName := fmt.Sprintf("../../internal/userfiles/converted_files/%s.%s", FileWithoutExt, outputformat) //This will create a file in that location
	cmd := exec.Command("ffmpeg", "-y", "-i", inputfilePath, "-c:v", "h264_nvenc", outputFileName)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(os.Getwd())
	if err != nil {
		fmt.Println("Error converting video:", err)
		return ""
	}

	fmt.Println("Video converted successfully!")
	return outputFileName //to get it to produce a message to kafka that it is converted
}

func VideoConversionConsumer() {

	c, _ := consumer.NewConsumer("broker:9092", "conversion_service")
	_ = c.Subscribe("conversion", nil)

	defer consumer.Close(c)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		var conversionMessage models.ConversionMessage
		err = json.Unmarshal(msg.Value, &conversionMessage)
		if err != nil {
			fmt.Println(err)
		}
		outputFilePath := Conversion(conversionMessage.FileName, conversionMessage.OutputFormat) //we get the file path here

		// Need to produce a message on the topic to upload
		//Prodcue the message for the mongo consumer with Video Key
		var messageToUpload models.AfterConvertUpload
		FileWithoutExt := strings.TrimSuffix(filepath.Base(conversionMessage.FileName), filepath.Ext(conversionMessage.FileName))
		fileNameCorrect := fmt.Sprintf("%s.%s", FileWithoutExt, conversionMessage.OutputFormat)

		messageToUpload.VideoKey = conversionMessage.VideoKey
		messageToUpload.FilePath = outputFilePath
		messageToUpload.UserId = conversionMessage.UserId
		messageToUpload.FileName = fileNameCorrect
		messageToUpload.UserName = conversionMessage.UserName
		serializedMessage, errorSerializing := json.Marshal(messageToUpload)
		if errorSerializing != nil {
			log.Println("Error serializing message", errorSerializing)

		}
		p, err := producer.NewProducer("broker:9092")
		producer.ProduceNewMessage(p, "upload", string(serializedMessage)) //Producing the converted mesaage
		//Now need the upload consumer
		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}
