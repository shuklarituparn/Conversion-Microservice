package email

import (
	"encoding/json"
	"fmt"
	"github.com/resend/resend-go/v2"
	"github.com/shuklarituparn/Conversion-Microservice/internal/consumer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"log"
	"os"
)

func SendMail(Filepath string, To string, Subject string) {

	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	emailToUse := fmt.Sprintf("Video Conversion Service <%s>", os.Getenv("EMAIL"))
	replyMail := os.Getenv("REPLY_MAIL")

	htmlcontent, err := os.ReadFile(Filepath)
	if err != nil {

		log.Println("Error opening the html content")
	}

	params := &resend.SendEmailRequest{
		From:    emailToUse,
		To:      []string{To},
		Html:    string(htmlcontent),
		Subject: Subject,
		ReplyTo: replyMail,
	}

	_, err = client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func GenerateVerficationEmailConsumer() {

	c, _ := consumer.NewConsumer("broker:9092", "email_service")
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
		sendMail.Subject = "Проверьте вашу почту"

		serializedMessage, err := json.Marshal(sendMail)
		if err != nil {
			log.Println("Failed to Serialize the message")
		}

		p, err := producer.NewProducer("broker:9092")
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

func GenerateRestoreEmailConsumer() {

	c, _ := consumer.NewConsumer("broker:9092", "email_service")
	_ = c.Subscribe("restore_mail", nil)

	defer consumer.Close(c)
	var RestoreMail models.RestoreAccountMessage

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		err = json.Unmarshal(msg.Value, &RestoreMail)
		filepath := RestoreIDTempGenerator(RestoreMail.UserName, RestoreMail.UserId)

		var sendmail models.MailSendMessage
		db := user_database.ReturnDbInstance()
		var user *models.User

		user, errorGettingUser := user_database.GetDeletedUserWithID(db, RestoreMail.UserId)
		if errorGettingUser != nil {
			log.Println("Error getting user from the database", errorGettingUser)
		}
		sendmail.TO = user.UserEmail
		sendmail.Filepath = filepath
		sendmail.Subject = "Удаление Аккаунта"

		serializedMessage, err := json.Marshal(sendmail)
		if err != nil {
			log.Println("Failed to Serialize the message")
		}

		p, err := producer.NewProducer("broker:9092")
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

func SendEmailConsumer() {

	c, _ := consumer.NewConsumer("broker:9092", "email_service")
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
		SendMail(mailToSend.Filepath, mailToSend.TO, mailToSend.Subject)

		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}

func DownloadMailConsumer() {

	c, _ := consumer.NewConsumer("broker:9092", "email_service")
	_ = c.Subscribe("download_mail", nil)

	defer consumer.Close(c)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		var FileDownloadMailMsg models.FiledownloadMailMessage
		err = json.Unmarshal(msg.Value, &FileDownloadMailMsg)
		if err != nil {
			fmt.Println(err)
		}
		pathOftemplate := FileDownloadTempGenerator(FileDownloadMailMsg.UserName, FileDownloadMailMsg.Mode, FileDownloadMailMsg.UserID, FileDownloadMailMsg.FileId)

		db := user_database.ReturnDbInstance() //getting the database instance

		//Need to get the userMail to send them the mail

		result, ErrorGettingUser := user_database.GetUserWithID(db, FileDownloadMailMsg.UserID)
		if ErrorGettingUser != nil {
			log.Println("Error getting UserMail in the DownloadMailGen", ErrorGettingUser)
		}

		var sendMailMsg models.MailSendMessage

		sendMailMsg.TO = result.UserEmail
		sendMailMsg.Filepath = pathOftemplate
		sendMailMsg.Subject = "Вот ваш файл"

		serialize, _ := json.Marshal(sendMailMsg)

		p, err := producer.NewProducer("broker:9092")

		producer.ProduceNewMessage(p, "send_mail", string(serialize))

		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}
