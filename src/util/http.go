package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, obj interface{}) {
	message, _ := json.Marshal(obj)
	w.Write(message)
}

func ResponseHTML(w http.ResponseWriter, html string) {
	w.Write([]byte(html))
}

func GetBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

func SendHTTPRequest(method string, url string, body []byte) error {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
		return err
	}
	defer response.Body.Close()

	return err
}