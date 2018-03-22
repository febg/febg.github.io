package api

import "github.com/gorilla/mux"

func GetRouter(c *Controller) *mux.Router {

	router := mux.NewRouter()
	router.Methods("GET").Path("/email/{client_email}/{client_name}/{client_number}").HandlerFunc(c.SendEmail)

	return router
}
