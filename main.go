package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("API_KEY")
	}
	if apiKey == "" {
		log.Fatal("RESEND_API_KEY or API_KEY is required in .env or environment")
	}

	toEnv := os.Getenv("RESEND_TO")
	if toEnv == "" {
		toEnv = os.Getenv("TO_EMAIL")
	}
	if toEnv == "" {
		log.Fatal("RESEND_TO or TO_EMAIL is required in .env or environment")
	}

	var toList []string
	for _, email := range strings.Split(toEnv, ",") {
		trimmed := strings.TrimSpace(email)
		if trimmed != "" {
			toList = append(toList, trimmed)
		}
	}
	if len(toList) == 0 {
		log.Fatal("No valid recipient email found in RESEND_TO")
	}

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      toList,
		Subject: "Hello World",
		Html:    "<p>Congrats on sending your <strong>first email</strong>!</p>",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		log.Fatalf("failed to send email: %v", err)
	}

	log.Printf("email sent: %s", sent.Id)
}
