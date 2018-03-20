package api

import (
	"log"
  "net/smtp"
)

type Controller struct {
	SmtpConf smtp.Auth
}

func NewController() (*Controller, error) {
	c := Controller{
		SmtpConf: smtp.PlainAuth(
        "",
        "felipe@cattivoragazzo.com",
        "1692Cattivo!",
        "smtp.gmail.com",
    ),
	}

	defer log.Printf("[INFO] Started email controller { [Hosted Rooms: %+v] }\n", c.SmtpConf)

	return &c, nil
}
