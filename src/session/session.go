package session

import (
	"os"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func InitStore() {
	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY"))) //mongodbstore.NewMongoDBStore(db.Session, []byte(os.Getenv("SESSION_KEY")))

}
