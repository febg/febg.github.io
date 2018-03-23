package api

import ("github.com/gorilla/mux"
        "net/http"
        "io"
        
        "gopkg.in/gomail.v2"
        )


// TODO: Implement gomail package in order to use subject, and attachments
func (c *Controller) SendEmail(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
    client_email := v["client_email"]
    client_name := v["client_name"]
    client_number := v["client_number"]
    password := v["password"]
    
    m := gomail.NewMessage()
    m.SetHeader("From", "felipe@cattivoragazzo.com")
    m.SetHeader("To", client_email)
    m.SetHeader("Subject", "Hello!")
    m.Embed("/Users/febg/Documents/CattivoRagazzoWeb/febg.github.io/Go/api/Legacy2.png")
    m.SetBody("text/html", `Thank you `+ client_name+`! <br /><br />

    Dino Minichiello and the team at Cattivo Ragazzo look forward to seeing you on April 5th, 2018, between 5pm-8pm for an outstanding event celebrating bespoke traditional and contemporary edge.<br /><br />
    
    Further event details will be sent via the registered email. <br /><br />
    
    If you have any questions, please do not hesitate to contact event coordinator Helen Siwak at helen@cattivoragazzo.com. <br /><br />
    
    Ciao! <br /><br />
    
    *** <br /> <h2>Helen Siwak</h2> M: 778-847-3011<br />Studio: 608-36 Water St, Vancouver BC <br /><b>By Appointment Only</b><br /> W: www.cattivoragazzo.com <br /> <img src="cid:Legacy2.png" alt="My image" /> <br />`+client_name+client_number+password)
 

    d := gomail.NewPlainDialer("smtp.gmail.com", 587, "felipe@cattivoragazzo.com", "1692Cattivo!")

    // Send the email to Bob, Cora and Dan.
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }


   

//   err := smtp.SendMail(
//       "smtp.gmail.com:587",
//       c.SmtpConf,
//       "felipe@cattivoragazzo.com",
//       []string{client_email},
//       []byte("This is the email body."+client_email+client_name+client_number+password),
//   )
//   if err != nil {
//       log.Fatal(err)
//   }

}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "ok")
}
