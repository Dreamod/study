package verify

import (
	"go/study/configs"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(vEmail string, hash string) error {
	conf := configs.LoadConfig()
	e := email.NewEmail()
	e.From = "Dreamod Go App <dreamod1@gmail.com>"
	e.To = []string{vEmail}
	e.Subject = "Подтверждение email"
	e.HTML = []byte("<a href=\"http://localhost:8081/verify/" + hash + "\">Подтвердить</a>")

	err := e.Send(conf.Email.Address+":587",
		smtp.PlainAuth("",
			conf.Email.Email,
			conf.Email.Password,
			conf.Email.Address))

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
