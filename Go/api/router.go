package api

import "github.com/gorilla/mux"

func GetRouter(c *Controller) *mux.Router {

	router := mux.NewRouter()

	//for a user to join a specific room
	router.Methods("GET").Path("/email/{client_email}/{client_name}/").HandlerFunc(c.SendEmail)

	return router
}
