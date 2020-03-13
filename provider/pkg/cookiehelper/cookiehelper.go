package cookiehelper

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func SetCookie(key string, value string) (res *http.Cookie) {
	cookieValue := map[string]string{
		"status": value,
	}
	if encoded, err := cookieHandler.Encode(key, cookieValue); err == nil {
		cookie := &http.Cookie{
			Name:   key,
			Value:  encoded,
			Secure: true,
			Domain: ".same-site-provider.herokuapp.com",
		}
		return cookie
	}
	return nil
}

func SetCookieWithSameSite(sameSiteMode http.SameSite, key string, value string) (res *http.Cookie) {
	cookie := SetCookie(key, value)
	cookie.SameSite = sameSiteMode
	return cookie
}

func GetValueByKey(key string, request *http.Request) (res string) {
	cookie, err := request.Cookie(key)
	if err != nil {
		return ""
	}

	cookieValue := make(map[string]string)
	err = cookieHandler.Decode(key, cookie.Value, &cookieValue)
	if err != nil {
		return ""
	}

	return cookieValue["status"]

}
