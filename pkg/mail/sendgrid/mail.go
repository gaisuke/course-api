package mail

import (
	"bytes"
	registerDto "course-api/internal/register/dto"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail interface {
	SendVerification(toEmail string, data registerDto.EmailVerification)
}

type mailUsecase struct {
}

// SendVerification implements Mail.
func (usecase *mailUsecase) SendVerification(toEmail string, data registerDto.EmailVerification) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/emails/verification_email.html")

	result, err := ParseTemplate(templateFile, data)
	if err != nil {
		fmt.Println(err)
	} else {
		usecase.sendMail(toEmail, data.SUBJECT, result)
	}
}

func (usecase *mailUsecase) sendMail(toEmail, subject, result string) {
	from := mail.NewEmail(os.Getenv("MAIL_SENDER_NAME"), os.Getenv("MAIL_SENDER_EMAIL"))
	to := mail.NewEmail(toEmail, toEmail)

	message := mail.NewSingleEmail(from, subject, to, "", result)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	resp, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp)
	} else {
		fmt.Println("Email sent successfully to %s", toEmail)
	}
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewMailUsecase() Mail {
	return &mailUsecase{}
}
