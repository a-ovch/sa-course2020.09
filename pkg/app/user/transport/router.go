package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/a-ovch/sa-course2020.09/pkg/app/user"
	"github.com/a-ovch/sa-course2020.09/pkg/user/app"
)

type Router struct {
	app *user.App
}

func (rt *Router) PostUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var ud app.UserData
	err := decoder.Decode(&ud)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	u, err := rt.app.AppService.CreateUser(&ud)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	_, _ = fmt.Fprintf(w, "{\"id\": \"%s\"}", u.ID)
}

func (rt *Router) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	u, err := rt.app.AppService.FindUser(uid)
	if err == app.ErrUserNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	encodedUser, err := json.Marshal(u)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	_, _ = fmt.Fprint(w, string(encodedUser))
}

func (rt *Router) DeleteUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "DeleteUser")
}

func (rt *Router) PutUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "PutUser")
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
}

func writeBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
}

func NewRouter(app *user.App) *Router {
	return &Router{app: app}
}
