package beforeupdate

import (
	"fmt"
	"net/http"

	cookiehelper "github.com/selcukusta/chromium-same-site-update/provider/pkg/cookiehelper"
	filehelper "github.com/selcukusta/chromium-same-site-update/provider/pkg/filehelper"
)

func GetPollPageHandler(response http.ResponseWriter, request *http.Request) {
	var bodyQuestions, _ = filehelper.LoadFile("/go/bin/templates/get-poll-questions-2.html")
	var bodyAlready, _ = filehelper.LoadFile("/go/bin/templates/get-poll-already.html")
	response.Header().Add("x-frame-options", "allow-from https://same-site-consumer.herokuapp.com")
	response.Header().Add("content-security-policy", "frame-ancestors 'self' https://same-site-consumer.herokuapp.com")
	response.Header().Add("Content-Type", "text/html")

	isPollActive := cookiehelper.GetValueByKey("poll_status_v2", request)
	if isPollActive != "" {
		fmt.Fprintf(response, bodyAlready, isPollActive)
	} else {
		fmt.Fprintf(response, bodyQuestions)
	}
}

func SubmitPollPageHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("x-frame-options", "allow-from https://same-site-consumer.herokuapp.com")
	response.Header().Add("content-security-policy", "frame-ancestors 'self' https://same-site-consumer.herokuapp.com")
	http.SetCookie(response, cookiehelper.SetCookie("poll_status_v2", request.FormValue("fav_lang")))
	http.Redirect(response, request, "/before-update/poll", 302)
}
