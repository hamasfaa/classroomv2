package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("SECRET-KEY"))

func GetSession(r *http.Request, key string) (*sessions.Session, error) {
	return store.Get(r, key)
}

func SaveSession(w http.ResponseWriter, r *http.Request, session *sessions.Session) error {
	return session.Save(r, w)
}
