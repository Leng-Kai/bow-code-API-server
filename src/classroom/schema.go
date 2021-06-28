package classroom

import . "github.com/Leng-Kai/bow-code-API-server/schemas"

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
	} `json:"creator"`
	Members []struct {
		Uid  string `json:"uid"`
		Name string `json:"name"`
		Role Role   `json:"role"`
	} `json:"members"`
	CoursePlan CoursePlan `json:"course_plan"`
	Grades     []struct {
		Event Test `json:"event"`
	} `json:"grades"`
}

type Test struct {
}
