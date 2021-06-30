package user

import (
	"fmt"
)

type authFunc func(interface{}) (string, error)

var authHandler = map[string]authFunc{
	"google":	googleAuthenticator,
	"facebook":	facebookAuthenticator,
	"twitter": 	twitterAuthenticator,
}

func globalUid(method string, auth_uid string) string {
	return fmt.Sprintf("%s_%s", method, auth_uid)
}

func googleAuthenticator(payload interface{}) (string, error) {
	return "", nil
}

func facebookAuthenticator(payload interface{}) (string, error) {
	return "", nil
}

func twitterAuthenticator(payload interface{}) (string, error) {
	return "", nil
}