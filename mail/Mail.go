package mail

import (
	"bytes"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/notgman/go-price/database"
	"github.com/notgman/go-price/models"
)

func getMessageString(from, to, subject, body string) []byte {
	headers := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n"

	return []byte(headers + body)
}

func parseTemplate(templateName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		log.Fatalln("Could not parse template", err)
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SendMail(product models.ProductScrape, price int) {
	email := database.GetUserMail(product.ID)

	SENDER, exists := os.LookupEnv("SENDER")
	if !exists {
		log.Fatalln("SENDER not found")
	}
	PASS, exists := os.LookupEnv("APP_PASSWORD")
	if !exists {
		log.Fatalln("APP_PASSWORD not found")
	}

	from := SENDER
	appPass := PASS
	to := []string{
		email,
	}
	s := "smtp.gmail.com:587"

	data := models.EmailData{
		Name:  product.Name,
		Price: price,
	}

	ms, err := parseTemplate("index.html", data)
	if err != nil {
		log.Fatalln("Could not parse template file", err)
	}

	b := getMessageString(from, to[0], "Price reduced!!!", ms)
	auth := smtp.PlainAuth("", from, appPass, "smtp.gmail.com")
	err = smtp.SendMail(s, auth, from, to, b)
	if err != nil {
		log.Fatalln("Could not send email", err)
	}
	log.Println("Mail sent")
}
