package email

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ID"
	"github.com/shuklarituparn/Conversion-Microservice/internal/consumer"
	gomail "gopkg.in/mail.v2"
	"log"
	"os"
	"path/filepath"
)

func SendMail() {
	m := gomail.NewMessage()

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	From := os.Getenv("EMAIL")
	Key := os.Getenv("EMAIL_KEY")
	m.SetHeader("From", From)
	m.SetHeader("To", "shukla.r@phystech.edu")
	m.SetHeader("Subject", "Imp mail")
	m.SetHeader("Message-ID", fmt.Sprintf("<%s@example.com>", ID.ReturnID()))

	// Get the current working directory
	currentPath, _ := os.Getwd()
	// Construct the file destination
	fileDestination := filepath.Join(currentPath, "../"+"../"+"internal"+"/"+"email"+"/"+"templates", "email.html")

	// Read the HTML content from the file
	htmlContent, err := os.ReadFile(fileDestination)
	if err != nil {
		fmt.Println("Error reading HTML content:", err)
		return
	}

	content := string(htmlContent)
	m.SetBody("text/html", content)

	d := gomail.NewDialer("smtp.yandex.com", 465, From, Key)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func ConsumeEmail() {

	c, _ := consumer.NewConsumer("localhost:9092", "email_service")
	_ = c.Subscribe("email", nil)

	defer consumer.Close(c)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		//EmailTempGenerator()
		SendMail()

		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}

//package email
//
//import (
//	"github.com/confluentinc/confluent-kafka-go/kafka"
//	"github.com/joho/godotenv"
//	gomail "gopkg.in/mail.v2"
//	"log"
//	"os"
//)
//
//func SendMail() {
//	m := gomail.NewMessage()
//
//	err := godotenv.Load("../../.env")
//	if err != nil {
//		log.Println("Error loading .env file:", err)
//		return
//	}
//
//	From := os.Getenv("EMAIL")
//	Key := os.Getenv("EMAIL_KEY")
//	m.SetHeader("From", From)
//	m.SetHeader("To", "rtprnshukla@gmail.com")
//	m.SetHeader("Subject", "Imp mail")
//	htmlContent, err := os.ReadFile("./preview.html")
//	if err != nil {
//		log.Println("Error reading preview.html:", err)
//		return
//	}
//	content := string(htmlContent)
//	m.SetBody("text/html", content)
//
//	d := gomail.NewDialer("smtp.yandex.com", 465, From, Key)
//
//	if err := d.DialAndSend(m); err != nil {
//		log.Println("Error sending email:", err)
//	}
//}
//
//var (
//	consumer *kafka.Consumer
//)
//
//func ConsumeEmail() {
//	c, err := kafka.NewConsumer(&kafka.ConfigMap{
//		"bootstrap.servers": "localhost:9092",
//		"group.id":          "email-consumer-group",
//		"auto.offset.reset": "earliest",
//	})
//	if err != nil {
//		log.Println("Error creating Kafka consumer:", err)
//		return
//	}
//	consumer = c
//
//	defer func(consumer *kafka.Consumer) {
//		err := consumer.Close()
//		if err != nil {
//
//		}
//	}(consumer)
//
//	err = consumer.SubscribeTopics([]string{"email"}, nil)
//	if err != nil {
//		log.Println("Error subscribing to Kafka topic:", err)
//		c.Close()
//		return
//	}
//	for {
//		msg, err := consumer.ReadMessage(-1)
//		if err != nil {
//			log.Println("Error reading message from Kafka:", err)
//			continue
//		}
//		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))
//		_, commitErr := consumer.CommitMessage(msg)
//		if commitErr != nil {
//			log.Printf("Failed to commit offset: %v", commitErr)
//		}
//	}
//}
