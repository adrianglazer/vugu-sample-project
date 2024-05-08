package auth

import (
	"fmt"
	"log"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/vugu/vugu/js"
)

const TestJWT = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE3MTUxOTE4OTMsImV4cCI6MTc0NjcyNzg5MywiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.I1HxLZDYF-6XDOn17QT1LOki2Z7qn8lJJvDjvmqJsmk"

type Auth struct {
	token *jwt.Token
}

func (a *Auth) IsAuthenticated() bool {
	storageToken := js.Global().Get("localStorage").Call("getItem", "token").String()

	// token, err := jwt.Parse(storageToken, func(t *jwt.Token) (interface{}, error) { return nil, nil })
	token, err := jwt.Parse(storageToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := []byte("qwertyuiopasdfghjklzxcvbnm123456")
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		log.Println(err)
		return false
	}

	a.token = token

	return true
}

func (a *Auth) SetJWT(token string) {
	js.Global().Get("localStorage").Call("setItem", "token", token)
}

func (a *Auth) RemoveJWT() {
	js.Global().Get("localStorage").Call("removeItem", "token")
}

type AuthRef struct{ *Auth }

type AuthSetter interface {
	SetAuth(a *Auth)
}

func (ar *AuthRef) SetAuth(a *Auth) { ar.Auth = a }
