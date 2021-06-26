package main

import (
	"fmt"
	"net/http"

	"github.com/J-HowHuang/bow-code/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	u := user.User{
		Uid:  "123",
		Name: "123",
	}
	fmt.Print(u)
	http.ListenAndServe(":8080", r)
}
