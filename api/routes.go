package api

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/alphagov/battlestations/github"
)

func logRequest(
	handler func(http.ResponseWriter, *http.Request),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL.String())
		handler(w, r)
	}
}

func withUser(
	store sessions.Store,
	handler func(github.User, http.ResponseWriter, *http.Request),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "battlestations")
		if user, ok := session.Values["user"]; ok {
			handler(user.(github.User), w, r)
		} else {
			http.Redirect(w, r, "/authorize", http.StatusFound)
		}
	}
}

func MakeRouter(authKey []byte, encKey []byte, githubService github.Service) *mux.Router {
	store := sessions.NewFilesystemStore("", authKey, encKey)
	r := mux.NewRouter()

	store.MaxLength(0)
	gob.Register(github.User{})

	r.HandleFunc("/",
		logRequest(
			withUser(store, func(user github.User, w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello %s!", *user.Details.Login)
			})))

	r.HandleFunc("/authorize",
		logRequest(
			func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, githubService.AuthURL(), http.StatusFound)
			}))

	r.HandleFunc("/authorized",
		logRequest(
			func(w http.ResponseWriter, r *http.Request) {
				var user github.User
				var err error

				if user, err = githubService.UserFromCode(r.FormValue("code")); err != nil {
					fmt.Fprintf(w, "ERROR: %s", err)
					return
				}

				session, _ := store.New(r, "battlestations")
				session.Values["user"] = user
				session.Save(r, w)

				http.Redirect(w, r, "/", http.StatusFound)
			}))

	return r
}
