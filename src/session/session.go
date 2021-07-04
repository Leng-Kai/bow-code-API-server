package session

import (
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func InitStore() {
	Store = sessions.NewCookieStore([]byte("secret-key")) //mongodbstore.NewMongoDBStore(db.Session, []byte(os.Getenv("SESSION_KEY")))

}
