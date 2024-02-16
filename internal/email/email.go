package email

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ID"
	"github.com/shuklarituparn/Conversion-Microservice/internal/consumer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
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

	currentPath, _ := os.Getwd()
	fileDestination := filepath.Join(currentPath, "../"+"../"+"internal"+"/"+"email"+"/"+"templates", "email.html")

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

func SendEmail() {

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

func GenerateEmail() {

	c, _ := consumer.NewConsumer("localhost:9092", "email_service")
	_ = c.Subscribe("verification_mail", nil)

	defer consumer.Close(c)
	var EmailMessage models.EmailVerificationMessage

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		err = json.Unmarshal(msg.Value, &EmailMessage)
		VerificationTempGenerator(EmailMessage.UserName, EmailMessage.UserID, EmailMessage.VerificationCode)
		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}
