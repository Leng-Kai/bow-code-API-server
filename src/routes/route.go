package routes

import (
	"net/http"

	"github.com/Leng-Kai/bow-code-API-server/course"
	"github.com/Leng-Kai/bow-code-API-server/user"
	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("POST", "/user", user.Register, nil)
	register("GET", "/user", user.GetUsers, nil)
	register("GET", "/user/{id}", user.GetUserByID, nil)
	register("PUT", "/user/{id}", user.UpdateUserByID, nil)
	/***************** Course *****************/
	register("GET", "/course", course.GetAll, nil)
	register("POST", "/course", course.CreateNew, nil)
	register("GET", "/course/{id}", course.GetCourseByID, nil)
	register("POST", "/course/{id}/block", course.CreateBlock, nil)
	register("GET", "/course/{id}/block", course.GetBlocksTitle, nil)
	/******************************************/
	register("GET", "/", healthyCheck, nil)
}

func healthyCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.HandleFunc(route.Pattern, route.Handler).Methods(route.Method)

		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
