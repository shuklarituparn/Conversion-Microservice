package email

import (
	"fmt"
	"github.com/matcornic/hermes/v2"
	"os"
	"path/filepath"
)

func WelcomeTempGenerator(userName string, userID int) {

	h := hermes.Hermes{

		Theme: &hermes.Default{},
		Product: hermes.Product{

			Name:      "Сервис конвертации видео",
			Link:      "https://knowing-gannet-actively.ngrok-free.app/",
			Logo:      "https://iili.io/J0hcSs4.png",
			Copyright: "© 2024 Сервис конвертации видео. Все права защищены",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: userName,
			Intros: []string{
				"Добро пожаловать в сервис конвертации видео, мы очень рады, что вы с нами.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Чтобы начать пользоваться нашим сервисом, пожалуйста, нажмите здесь:",
					Button: hermes.Button{
						Color: "#0077FF",
						Text:  "Подтвердите свой адрес электронной почты",
						Link:  "https://knowing-gannet-actively.ngrok-free.app/",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	_, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	currentWorkDir, _ := os.Getwd()
	finalFilePath := filepath.Join(currentWorkDir, "templates")
	filename := fmt.Sprintf("%d_w.html", userID)
	filePath := filepath.Join(finalFilePath, filename)

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = os.WriteFile(finalFilePath+"/"+filename, []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}
}

func VerificationTempGenerator(userName string, userID int, VerificationCode string) string {

	userWelcomeString := fmt.Sprintf("Добро пожаловать в сервис конвертации видео, мы очень рады, что вы с нами.")
	userEmailString := fmt.Sprintf("https://knowing-gannet-actively.ngrok-free.app/verify_mail?code=%s&userId=%d", VerificationCode, userID)

	h := hermes.Hermes{

		Theme: &hermes.Default{},
		Product: hermes.Product{

			Name:        "Сервис конвертации видео",
			Link:        "https://knowing-gannet-actively.ngrok-free.app/",
			Logo:        "https://iili.io/J0hcSs4.png",
			Copyright:   "© 2024 Сервис конвертации видео. Все права защищены",
			TroubleText: "Если у вас возникли проблемы с кнопкой '{ACTION}', скопируйте и вставьте приведенный ниже URL-адрес в свой веб-браузер.",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: userName,
			Intros: []string{
				userWelcomeString,
			},
			Signature: "C Уважением",
			Greeting:  "Привет",
			Actions: []hermes.Action{
				{
					Instructions: "Чтобы подтвердить свой адрес электронной почты, пожалуйста, нажмите здесь:",
					Button: hermes.Button{
						Color: "#0077FF",
						Text:  "Подтвердить почту",
						Link:  userEmailString,
					},
				},
			},
			Outros: []string{
				"Если ты столкнулся с проблемой, ответь на это письмо, и мы поможем тебе прямо сейчас!",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	_, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	filename := fmt.Sprintf("%d_w.html", userID)
	filePath := filepath.Join("../../internal/email/templates", filename)

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	completeFilename := fmt.Sprintf("../../internal/email/templates" + "/" + filename)
	err = os.WriteFile(completeFilename, []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}
	return completeFilename
}

func RestoreIDTempGenerator(userName string, userID int) string {

	userWelcomeString := fmt.Sprintf("Мы очень рады, что вы выбрали нас! Мы надеемся, что встретимся снова!.")
	userEmailString := fmt.Sprintf("https://knowing-gannet-actively.ngrok-free.app/profile/restore")

	h := hermes.Hermes{

		Theme: &hermes.Default{},
		Product: hermes.Product{

			Name:        "Сервис конвертации видео",
			Link:        "https://knowing-gannet-actively.ngrok-free.app/",
			Logo:        "https://iili.io/J0hcSs4.png",
			Copyright:   "© 2024 Сервис конвертации видео. Все права защищены",
			TroubleText: "Если у вас возникли проблемы с кнопкой '{ACTION}', скопируйте и вставьте приведенный ниже URL-адрес в свой веб-браузер.",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: userName,
			Intros: []string{
				userWelcomeString,
			},
			Signature: "C Уважением",
			Greeting:  "До свидания",
			Actions: []hermes.Action{
				{
					Instructions: "Чтобы восстановить ваши данные, нажмите здесь:",
					Button: hermes.Button{
						Color: "#0077FF",
						Text:  "восстановить данные",
						Link:  userEmailString,
					},
				},
			},
			Outros: []string{
				"Если ты столкнулся с проблемой, ответь на это письмо, и мы поможем тебе прямо сейчас!",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	_, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	filename := fmt.Sprintf("%d_bye.html", userID)
	filePath := filepath.Join("../../internal/email/templates", filename)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	completeFilename := fmt.Sprintf("../../internal/email/templates" + "/" + filename)
	err = os.WriteFile(completeFilename, []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}
	return completeFilename
}

func FileDownloadTempGenerator(userName string, mode string, userID int, fileId string) string {

	userWelcomeString := fmt.Sprintf("Мы очень рады, что вы выбрали нас! Вы хотели %s ваши файли!. Вот вам ссылка чтобы загрузить ваш файл", mode)
	userEmailString := fmt.Sprintf("https://knowing-gannet-actively.ngrok-free.app/profile/download?userid=%d&fileid=%s&mode=%s", userID, fileId, mode)

	h := hermes.Hermes{

		Theme: &hermes.Default{},
		Product: hermes.Product{

			Name:        "Сервис конвертации видео",
			Link:        "https://knowing-gannet-actively.ngrok-free.app/",
			Logo:        "https://iili.io/J0hcSs4.png",
			Copyright:   "© 2024 Сервис конвертации видео. Все права защищены",
			TroubleText: "Если у вас возникли проблемы с кнопкой '{ACTION}', скопируйте и вставьте приведенный ниже URL-адрес в свой веб-браузер.",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: userName,
			Intros: []string{
				userWelcomeString,
			},
			Signature: "C Уважением",
			Greeting:  "Привет",
			Actions: []hermes.Action{
				{
					Instructions: "Чтобы загрузить ваш файл, нажмите здесь:",
					Button: hermes.Button{
						Color: "#0077FF",
						Text:  "Загруззить",
						Link:  userEmailString,
					},
				},
			},
			Outros: []string{
				"Если ты столкнулся с проблемой, ответь на это письмо, и мы поможем тебе прямо сейчас!",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	_, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	filename := fmt.Sprintf("%d_%s_mail.html", userID, mode)
	filePath := filepath.Join("../../internal/email/templates", filename)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	completeFilename := fmt.Sprintf("../../internal/email/templates" + "/" + filename)
	err = os.WriteFile(completeFilename, []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}
	return completeFilename
} //A single template generator for all type of file. Using the mode and the file name and userID
