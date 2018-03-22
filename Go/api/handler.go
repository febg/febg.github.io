package api

import ("github.com/gorilla/mux"
        "net/http"
        "net/smtp"
        "log")


func (c *Controller) SendEmail(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
    client_email := v["client_email"]
    client_name := v["client_name"]
    client_number := v["client_number"]
    
  err := smtp.SendMail(
      "smtp.gmail.com:587",
      c.SmtpConf,
      "felipe@cattivoragazzo.com",
      []string{client_email},
      []byte("This is the email body."),
  )
  if err != nil {
      log.Fatal(err)
  }

}
