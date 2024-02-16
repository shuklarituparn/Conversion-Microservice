package email

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shuklarituparn/Conversion-Microservice/internal/consumer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	gomail "gopkg.in/mail.v2"
	"log"
	"os"
)

func SendMail(Filepath string, To string) {
	m := gomail.NewMessage()

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	From := os.Getenv("EMAIL")
	Key := os.Getenv("EMAIL_KEY")
	m.SetHeader("From", From)
	m.SetHeader("To", To)
	m.SetHeader("Subject", "Verification Mail")

	htmlContent, err := os.ReadFile(Filepath)
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

func GenerateVerficationEmail() {

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
		filepath := VerificationTempGenerator(EmailMessage.UserName, EmailMessage.UserID, EmailMessage.VerificationCode)

		var sendMail models.MailSendMessage

		sendMail.Filepath = filepath
		sendMail.TO = EmailMessage.UserEmail

		serializedMessage, err := json.Marshal(sendMail)
		if err != nil {
			log.Println("Failed to Serialize the message")
		}

		p, err := producer.NewProducer("localhost:9092")
		err = producer.ProduceNewMessage(p, "send_mail", string(serializedMessage))
		if err != nil {
			return
		}

		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}

func SendEmail() {

	c, _ := consumer.NewConsumer("localhost:9092", "email_service")
	_ = c.Subscribe("send_mail", nil)

	defer consumer.Close(c)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		var mailToSend models.MailSendMessage
		err = json.Unmarshal(msg.Value, &mailToSend)
		if err != nil {
			fmt.Println(err)
		}
		SendMail(mailToSend.Filepath, mailToSend.TO)

		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}
