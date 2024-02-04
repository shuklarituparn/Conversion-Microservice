package email

import (
	"crypto/tls"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ID"
	"github.com/shuklarituparn/Conversion-Microservice/internal/consumer"
	gomail "gopkg.in/mail.v2"
	"log"
	"os"
)

func SendMail() {
	m := gomail.NewMessage()

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	From := os.Getenv("EMAIL")
	Key := os.Getenv("EMAIL_KEY")

	// Set E-Mail sender
	m.SetHeader("From", From)

	// Set E-Mail receivers
	m.SetHeader("To", "shukla.r@phystech.edu")

	// Set E-Mail subject
	m.SetHeader("Subject", "Imp mail")

	m.SetHeader("Message-ID", fmt.Sprintf("<%s@example.com>", ID.ReturnID()))

	htmlContent, err := os.ReadFile("preview.html")
	content := string(htmlContent)
	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", content)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.yandex.com", 465, From, Key)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
func ConsumeEmail() {

	c, err := consumer.NewConsumer("localhost:9092", "email", []string{"upload"})
	if err != nil {
		log.Println("Error creating consumer:", err)
		return
	}
	defer c.Close()

	for {

		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		SendMail()

		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}
}
