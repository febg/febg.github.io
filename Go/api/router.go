package api

import "github.com/gorilla/mux"

func GetRouter(c *Controller) *mux.Router {

	router := mux.NewRouter()
	router.Methods("GET").Path("/email/{client_email}/{client_name}/{client_number}/{password}").HandlerFunc(c.SendEmail)
	router.Methods("GET").Path("/ping").HandlerFunc(c.Ping)

	return router
}
