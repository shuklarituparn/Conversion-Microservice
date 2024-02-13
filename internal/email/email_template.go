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
	err = os.WriteFile(finalFilePath+"/"+filename, []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}
}

func TempGenerator(userName string, userID int) {

	h := hermes.Hermes{

		Theme: &hermes.Default{},
		Product: hermes.Product{

			Name:      "Video conversion service",
			Link:      "https://knowing-gannet-actively.ngrok-free.app/",
			Logo:      "https://iili.io/J0hcSs4.png",
			Copyright: "Copyright © 2024 Video Conversion Service. All rights reserved",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: userName,
			Intros: []string{
				"Welcome to Video Conversion Service! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Video Conversion Service, please click here:",
					Button: hermes.Button{
						Color: "#0077FF",
						Text:  "Confirm your account",
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
	err = os.WriteFile(finalFilePath+"/"+filename, []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}
}
