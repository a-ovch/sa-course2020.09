package transport

import (
	"fmt"
	"github.com/a-ovch/sa-course2020.09/pkg/app/user"
	"net/http"
)

type Router struct {
	app *user.App
}

func (r *Router) PostUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "PostUser")
}

func (r *Router) GetUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "GetUser")
}

func (r *Router) DeleteUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "deleteUserHandler")
}

func (r *Router) PutUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "putUserHandler")
}

func NewRouter(app *user.App) *Router {
	return &Router{app: app}
}
