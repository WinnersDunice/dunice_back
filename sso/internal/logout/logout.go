package logout

import (
	"net/http"

	ts "github.com/WinnersDunice/dunice_back/sso/internal/tokenset"
	"github.com/gorilla/sessions"
)

func LogoutHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "auth")

		ts.Delete(session)
		session.Options.MaxAge = -1
		session.Save(r, w)

	}
}
