package services

import (
	"os"

	gomail "github.com/go-mail/mail"
)

// SendAlert : Send an email the address in config with the given message
func SendAlert(email, message string) {
	d := gomail.NewDialer("zimbra.enix.io", 25, "achaloin", os.Getenv("EMAIL_PASSWORD"))
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "alerts@banana.enix.io")
	m.SetHeader("To", email)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hellosssss!")
	m.SetBody("text/html", message)
	// m.Attach("/home/Alex/lolcat.jpg")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
