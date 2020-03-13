package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	common "github.com/selcukusta/chromium-same-site-update/consumer/common"
)

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", common.MainPageHandler)

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")), nil)
}
