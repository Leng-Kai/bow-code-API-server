package routes

import (
	"net/http"

	"github.com/Leng-Kai/bow-code-API-server/course"
	"github.com/Leng-Kai/bow-code-API-server/user"
	"github.com/Leng-Kai/bow-code-API-server/problem"
	"github.com/Leng-Kai/bow-code-API-server/submit"
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
	/* User */
	register("POST", "/register", user.Register, nil)
	register("POST", "/login", user.Login, nil)
	register("POST", "/logout", user.Logout, nil)
	register("GET", "/auth", user.AuthSession, nil)
	register("GET", "/user/{id}", user.GetUserByID, nil)
	register("PUT", "/user/{id}", user.UpdateUserByID, nil)

	/* Course */
	register("GET", "/course", course.GetAll, nil)
	register("POST", "/course", course.CreateNew, nil)
	register("GET", "/course/multiple", course.GetMultipleCourses, nil)
	register("GET", "/course/{id}", course.GetCourseByID, nil)
	register("PUT", "/course/{id}", course.UpdateCourseByID, nil)
	register("DELETE", "/course/{id}", course.RemoveCourseByID, nil)
	register("POST", "/course/{id}/block", course.CreateBlock, nil)
	register("GET", "/course/{id}/block/{bid}", course.GetBlock, nil)
	register("PUT", "/course/{id}/block/{bid}", course.UpdateBlock, nil)
	register("PUT", "/course/{id}/blockOrder", course.ModifyBlockOrder, nil)
	register("POST", "/course/{id}/favorite", course.LoveCourseByID, nil)

	/* Problem */
	register("GET", "/problem", problem.GetProblems, nil)
	register("POST", "/problem", problem.CreateNewProblem, nil)
	// register("GET", "/problem/multiple", problem.GetMultipleProblems, nil)
	register("GET", "/problem/{pid}", problem.GetProblemByID, nil)
	register("POST", "/problem/{pid}", problem.UpdateProblemByID, nil)
	
	/* Submit */
	register("GET", "/submit/multiple", submit.GetMultipleSubmissions, nil)
	register("GET", "/submit/user", submit.GetSubmissionsByUID, nil)
	register("GET", "/submit/{sid}", submit.GetSubmissionByID, nil)
	register("PUT", "/submit/{sid}/{caseNo}", submit.ReceiveJudgeResult, nil)
	register("POST", "/submit/problem/{pid}", submit.SubmitToProblem, nil)
	
	/* Healthy Check */
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
