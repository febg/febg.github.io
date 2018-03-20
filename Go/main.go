package main

import (
    "log"
    
    "./api"
    	"net/http"
)

func main() {
    // Set up authentication information.
    controller, err := api.NewController()


    router := api.GetRouter(controller)
    // Connect to the server, authenticate, set the sender and recipient,
    // and send the email all in one step.
    log.Println("[INFO] Listening on http://localhost:8081")
    err = http.ListenAndServe(":8081", router)
    if err != nil {
      log.Fatal("ListenAndServe Error: ", err)
    }
}
