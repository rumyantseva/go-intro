package main

import (
	"fmt"

	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/api/v1/users", createUser)
	router.GET("/api/v1/users/:id", findUserByID)

	http.ListenAndServe(":8000", router)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Some code to register user...\n")
}

func findUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Some code to get user...\n")
}
