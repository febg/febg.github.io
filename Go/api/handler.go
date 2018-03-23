package api

import ("github.com/gorilla/mux"
        "net/http"
        "io"
        "net/smtp"
        "log")


// TODO: Implement gomail package in order to use subject, and attachments
func (c *Controller) SendEmail(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
    client_email := v["client_email"]
    client_name := v["client_name"]
    client_number := v["client_number"]
    password := v["password"]
    
  err := smtp.SendMail(
      "smtp.gmail.com:587",
      c.SmtpConf,
      "felipe@cattivoragazzo.com",
      []string{client_email},
      []byte("This is the email body."+client_email+client_name+client_number+password),
  )
  if err != nil {
      log.Fatal(err)
  }

}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "ok")
}
