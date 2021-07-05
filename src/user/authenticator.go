package user

import (
	"fmt"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
)

type authFunc func(interface{}) (string, UserInfo, error)

var authHandler = map[string]authFunc{
	"google":	googleAuthenticator,
	"facebook":	facebookAuthenticator,
	"twitter": 	twitterAuthenticator,
}

func globalUid(method string, auth_uid string) string {
	return fmt.Sprintf("%s_%s", method, auth_uid)
}

func googleAuthenticator(payload interface{}) (string, UserInfo, error) {
	name := "George"
	avatar := "new_avatar_url"
	return "googleUidForTesting", UserInfo{name, avatar}, nil
}

func facebookAuthenticator(payload interface{}) (string, UserInfo, error) {
	name := "Frank"
	avatar := "avatar_url"
	return "facebookUidForTesting", UserInfo{name, avatar}, nil
}

func twitterAuthenticator(payload interface{}) (string, UserInfo, error) {
	name := "Taylor"
	avatar := "avatar_url"
	return "twitterUidForTesting", UserInfo{name, avatar}, nil
}