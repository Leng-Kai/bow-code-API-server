package classroom

import "github.com/J-HowHuang/bow-code/course"

type Role int

const (
	teacher Role = iota
	TA
	student
)

func (role Role) String() string {
	return [...]string{"teacher", "TA", "student"}[role]
}

type Classroom struct {
	ClassId string `json:"class_id"`
	Creator struct {
		Uid  string `json:"uid"`
		Name string `json:"name"`
	}`json:"creator"`
	Members []struct {
		Uid  string `json:"uid"`
		Name string `json:"name"`
		Role Role `json:"role"`
	}`json:"members"`
	CoursePlan course.CoursePlan `json:"course_plan"`
	Grades []struct{
		Event Test `json:"event"`
		Students
	} `json:"grades"`
}

type Test struct {

}