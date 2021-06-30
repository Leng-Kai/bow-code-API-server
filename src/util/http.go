package util

import (
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
