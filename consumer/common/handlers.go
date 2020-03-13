package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// for GET
func MainPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = LoadFile("/go/bin/templates/main.html")
	fmt.Fprintf(response, body)
}
