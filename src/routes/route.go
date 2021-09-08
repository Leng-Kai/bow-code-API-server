package routes

import (
	"net/http"

	"github.com/Leng-Kai/bow-code-API-server/bulletin"
	"github.com/Leng-Kai/bow-code-API-server/classroom"
	"github.com/Leng-Kai/bow-code-API-server/course"
	"github.com/Leng-Kai/bow-code-API-server/course_plan"
	"github.com/Leng-Kai/bow-code-API-server/submit"
	"github.com/Leng-Kai/bow-code-API-server/user"
	"github.com/Leng-Kai/bow-code-API-server/problem"
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
	register("GET", "/course/details", course.GetCoursesDetails, nil)
	register("GET", "/course/{id}", course.GetCourseByID, nil)
	register("PUT", "/course/{id}", course.UpdateCourseByID, nil)
	register("DELETE", "/course/{id}", course.RemoveCourseByID, nil)
	register("POST", "/course/{id}/block", course.CreateBlock, nil)
	register("PUT", "/course/{id}/block/title/{bid}", course.UpdateBlockTitle, nil)
	register("GET", "/course/{id}/block/{bid}", course.GetBlock, nil)
	register("PUT", "/course/{id}/block/{bid}", course.UpdateBlock, nil)
	register("PUT", "/course/{id}/blockOrder", course.ModifyBlockOrder, nil)
	register("POST", "/course/{id}/favorite", course.LoveCourseByID, nil)

	/* CoursePlan */
	register("GET", "/course_plan", course_plan.GetCoursePlans, nil)
	register("POST", "/course_plan", course_plan.CreateNewCoursePlan, nil)
	register("GET", "/course_plan/{cpid}", course_plan.GetCoursePlanByID, nil)
	register("POST", "/course_plan/{cpid}", course_plan.UpdateCoursePlanByID, nil)

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
	register("PUT", "/submit/{crid}/{sid}/{caseNo}", submit.ReceiveJudgeResult_Classroom, nil)
	register("POST", "/submit/problem/{pid}", submit.SubmitToProblem, nil)
	// register("POST", "/submit/{crid}/problem/{pid}", submit.SubmitToProblem_Classroom, nil)

	/* Classroom */
	register("GET", "/classroom", classroom.GetClassrooms, nil)
	register("POST", "/classroom", classroom.CreateNewClassroom, nil)
	register("POST", "/classroom/apply/{crid}", classroom.ApplyForClassroom, nil) 
	register("POST", "/classroom/accept/{crid}/{uid}", classroom.AcceptApplication, nil)
	register("POST", "/classroom/reject/{crid}/{uid}", classroom.RejectApplication, nil)
	register("POST", "/classroom/invite/{crid}/{uid}", classroom.InviteStudent, nil)
	register("POST", "/classroom/join/{crid}", classroom.JoinClassroom, nil)
	register("GET", "/classroom/status/{crid}", classroom.GetClassroomStatus, nil)
	register("GET", "/classroom/record/{crid}", classroom.GetClassroomRecord, nil)
	register("GET", "/classroom/score/{crid}/{uid}", classroom.GetStudentScores, nil)
	register("POST", "/classroom/homework/{crid}", classroom.CreateHomework, nil)
	register("POST", "/classroom/exam/{crid}", classroom.CreateExam, nil)
	register("PUT", "/classroom/homework/{crid}/{No}", classroom.UpdateHomework, nil)
	register("PUT", "/classroom/exam/{crid}/{No}", classroom.UpdateExam, nil)
	register("GET", "/classroom/bulletin/{crid}", classroom.GetAllBulletins, nil)
	register("GET", "/classroom/{crid}", classroom.GetClassroomByID, nil)
	register("POST", "/classroom/{crid}", classroom.UpdateClassroomByID, nil)

	register("GET", "/bulletin/{bid}", bulletin.GetBulletin, nil)
	register("POST", "/bulletin/{crid}", bulletin.CreateNewBulletin, nil)
	register("POST", "/bulletin/reply/like/{bid}/{index}", bulletin.LikeReply, nil)
	register("POST", "/bulletin/reply/unlike/{bid}/{index}", bulletin.UnlikeReply, nil)
	register("POST", "/bulletin/reply/{bid}", bulletin.ReplyToBulletin, nil)
	register("DELETE", "/bulletin/reply/{bid}/{index}", bulletin.DeleteReply, nil)
	register("PUT", "/bulletin/reply/{bid}/{index}", bulletin.EditReply, nil)
	register("POST", "/bulletin/like/{bid}", bulletin.LikeBulletin, nil)
	register("POST", "/bulletin/unlike/{bid}", bulletin.UnlikeBulletin, nil)
	register("DELETE", "/bulletin/{bid}", bulletin.DeleteBulletin, nil)
	register("PUT", "/bulletin/{bid}", bulletin.EditBulletin, nil)
	
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
