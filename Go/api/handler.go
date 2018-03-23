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
    // client_name := v["client_name"]
    // client_number := v["client_number"]
     password := v["password"]
    
    m := gomail.NewMessage()
    if (password == "Bespoke4Life"){
        m.SetBody("text/html", `<body style ="text-align:center">Thank you for RSVPing for our VIP Event celebrating bespoke suiting. <br /><br />

            You have been added to the guest list and will receive
            
            a follow-up email with additional details in the days before
            
            revealing the location and libation selection.<br /><br/>
            
            If you have any questions regarding the event, Cattivo Ragazzo, or the Master Tailors,
            please do not hesitate to contact helen@cattivoragazzo.com
            
           </body>`)
    }else{
        m.SetBody("text/html", `<body style ="text-align:center">Thank you for booking a consultation for our VIP Event celebrating bespoke suiting. <br /><br />

            Dino Minichiello and the Cattivo Ragazzo team look forward to welcoming you 
            
            to our atelier for your personalized appointment 
            
            with Master Tailors Verolo & Lemmi.<br /><br/>
            
            If you have any questions in advance of setting your appointment,
Cattivo Ragazzo, or the Master Tailors,
please do not hesitate to contact tania@cattivoragazzo.com.
            
           </body>`) }
   
    m.SetHeader("From", "Helen@cattivoragazzo.com")
    m.SetHeader("To", client_email)
    m.SetHeader("Subject", "Confirmation of Cattivo Ragazzo VIP event RSVP")
    
 

    d := gomail.NewPlainDialer("smtp.gmail.com", 587, "Helen@cattivoragazzo.com", "Victoria1122!!")

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
