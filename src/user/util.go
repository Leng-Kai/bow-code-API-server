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
	if islogin, ok := session.Values["isLogin"].(bool); !ok || !islogin {
		return schemas.User{}, errors.New("user not login")
	}
	id, ok := session.Values["uid"].(string)
	if !ok {
		return schemas.User{}, errors.New("invalid id")
	}
	user, err := db.GetSingleUserByID(id)
	if err != nil {
		return schemas.User{}, errors.New("user not exist")
	}
	return user, nil
}
