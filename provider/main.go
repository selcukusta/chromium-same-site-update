package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	afterupdate "github.com/selcukusta/chromium-same-site-update/provider/pkg/afterupdate"
	beforeupdate "github.com/selcukusta/chromium-same-site-update/provider/pkg/beforeupdate"
)

var router = mux.NewRouter()

func main() {

	router.HandleFunc("/after-update/poll", afterupdate.GetPollPageHandler)
	router.HandleFunc("/after-update/poll-submit", afterupdate.SubmitPollPageHandler).Methods("POST")

	router.HandleFunc("/before-update/poll", beforeupdate.GetPollPageHandler)
	router.HandleFunc("/before-update/poll-submit", beforeupdate.SubmitPollPageHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")), nil)
}
