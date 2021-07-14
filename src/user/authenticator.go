package user

import (
	"context"
	"fmt"
	"os"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/idtoken"
)

type authFunc func(interface{}) (string, UserInfo, error)

var authHandler = map[string]authFunc{
	"google":   googleAuthenticator,
	"facebook": facebookAuthenticator,
	"twitter":  twitterAuthenticator,
}

var api_client_id = map[string]string{}

func init() {
	api_client_id["google"] = os.Getenv("GOOGLE_CLIENT_ID")
}

func globalUid(method string, auth_uid string) string {
	return fmt.Sprintf("%s_%s", method, auth_uid)
}

func googleAuthenticator(token interface{}) (string, UserInfo, error) {

	idToken := token.(string)

	payload, err := idtoken.Validate(context.Background(), idToken, api_client_id["google"])
	if err != nil {
		return "", UserInfo{}, err
	}

	// some validity checks

	name := payload.Claims["name"].(string)
	avatar := payload.Claims["picture"].(string)
	uid := payload.Claims["sub"].(string)

	return uid, UserInfo{name, avatar}, nil
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
