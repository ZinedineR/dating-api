package mail

import (
	"dating-api/internal/profile/domain"
	"log"
	"os"
	"time"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "SMTP Zinedine <zinedin505@gmail.com>"
const CONFIG_AUTH_EMAIL = "zinedin505@gmail.com"
const CONFIG_AUTH_PASSWORD = "hvoipazrqtpdubev"

func GenerateWIB(t time.Time) time.Time {
	wib, err := time.LoadLocation("Asia/Jakarta")
	if err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), wib)
	}
	return t
}

func Verify_mail(form *domain.Profile) {
	var link, text string
	expires := form.CreatedAt.Add(time.Hour * 24)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", form.Email, "voucherlabs.official@gmail.com")
	mailer.SetAddressHeader("Cc", "voucherlabs.official@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Code Verification")
	link = os.Getenv("RUNNING_HOST") + "verify/" + form.Id.String()
	text = "Please, click this link before " + expires.Format(time.RFC1123)
	mailer.SetBody("text/html", "Hello, "+form.Name+"<br>have a nice day, here's your credentials : <br> Email : "+form.Email+"<br> Password : "+form.Password+"<br><br>"+text)
	mailer.SetBody("text/html", "Here's your verification link, click link<br><a href='"+link+"'>Click here</a><br>"+text)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Print(dialer)
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

func Forgot_password(form domain.Profile) {
	var link, text string
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", form.Email, "voucherlabs.official@gmail.com")
	mailer.SetAddressHeader("Cc", "voucherlabs.official@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Code Verification")
	link = os.Getenv("RUNNING_HOST") + "reset-password/" + form.Id.String()

	mailer.SetBody("text/html", "Please click this link to reset your password<br><a href='"+link+"'>Click here</a><br>"+text)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Print(dialer)
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
