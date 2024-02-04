package main

import (
	"github.com/matcornic/hermes/v2"
	"os"
)

func main() {
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		Theme: &hermes.Default{},
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Video conversion service",
			Link: "https://knowing-gannet-actively.ngrok-free.app/",
			// Optional product logo
			Logo:      "https://iili.io/J0hcSs4.png",
			Copyright: "Copyright © 2024 Video Conversion Service. All rights reserved",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "Rituparn Shukla",
			Intros: []string{
				"Welcome to Video Conversion Service! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Video Conversion Service, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
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

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	_, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Optionally, preview the generated HTML e-mail by writing it to a local file
	err = os.WriteFile("preview.html", []byte(emailBody), 0644)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}
}

//TODO: SO CAN BASICALLY GENERATE DIFFERENT TEMPLATES FOR DIFFERENT USE AND THEN USE KAFKA TO SEND TO USER
//TODO: Theme is where the general template is defined
