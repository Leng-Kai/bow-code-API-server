package user

import (
	"errors"
	"net/http"

	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/session"
)

func GetSessionUser(r *http.Request) (schemas.User, error) {
	session, err := session.Store.Get(r, "bow-session")
	if err != nil {
		return schemas.User{}, err
	}
	if !session.Values["isLogin"].(bool) {
		return schemas.User{}, errors.New("user not login")
	}
	id := session.Values["uid"]
	user, err := db.GetSingleUserByID(id.(string))
	if err != nil {
		return schemas.User{}, errors.New("user not exist")
	}
	return user, nil
}
