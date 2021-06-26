package util

import (
	"encoding/json"
	"net/http"
)

func responseJSON(w http.ResponseWriter, obj interface{}) {
	message, _ := json.Marshal(obj)
	w.Write(message)
}

func responseHTML(w http.ResponseWriter, html string) {
	w.Write([]byte(html))
}
