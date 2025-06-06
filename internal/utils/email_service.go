package utils

import (
	"fmt"
	"log"
	"mikhael-project-go/config"
	"mikhael-project-go/pkg/drivers/common"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "mikhael.wellman@idstar.group")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	env := common.NewEmailInfo(
		config.Config("SMTP"),
		config.Config("PORT_SMTP"),
		config.Config("EMAIL_SMTP"),
		config.Config("PASSWORD_SMTP"),
	)

	// d := gomail.NewDialer("smtp.gmail.com", 587, "mikhael.wellman@idstar.group", "aptb lwbi bhxb olbb")
	portStmp := env.Port
	port, _ := strconv.Atoi(portStmp)

	log.Println("haisl : ", env.Smtp, env.Port, env.Email, env.Password)
	d := gomail.NewDialer(env.Smtp, port, env.Email, env.Password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Gagal mengirim email:", err)
		return
	}

	fmt.Println("Email berhasil dikirim!")
}
