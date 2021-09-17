package schemas

type UserID = string

type UserInfo struct {
	Name   string `json:"name" bson:"name"`
	Avatar string `json:"avatar" bson:"avatar"`
	Email  string `json:"email" bson:"email"`
}

type SchoolInfo struct {
	SchoolName  string `json:"schoolName" bson:"schoolName"`
	StudentID   string `json:"studentID" bson:"studentID"`
	StudentName string `json:"studentName" bson:"studentName"`
	Class       string `json:"class" bson:"class"`
	SeatNumber  int    `json:"seatNumber" bson:"seatNumber"`
}

type User struct {
	UserID              UserID     `json:"id" bson:"_id,omitempty"`
	RegisterType        string     `json:"registerType" bson:"registerType"`
	UserInfo            UserInfo   `json:"userInfo" bson:"userInfo"`
	SchoolInfo          SchoolInfo `json:"schoolInfo" bson:"schoolInfo"`
	SchoolInfoCompleted bool       `json:"schoolInfoCompleted" bson:"schoolInfoCompleted"`
	Super               bool       `json:"super" bson:"super"`

	JoinedCourseList   []CourseID `json:"joinedCourseList" bson:"joinedCourseList"`
	OwnCourseList      []CourseID `json:"ownCourseList" bson:"ownCourseList"`
	FavoriteCourseList []CourseID `json:"favoriteCourseList" bson:"favoriteCourseList"`

	JoinedCoursePlanList []CoursePlanID `json:"joinedCoursePlanList" bson:"joinedCoursePlanList"`
	OwnCoursePlanList    []CoursePlanID `json:"ownCoursePlanList" bson:"ownCoursePlanList"`

	OwnProblemList      []ProblemID `json:"ownProblemList" bson:"ownProblemList"`
	FavoriteProblemList []ProblemID `json:"favoriteProblemList" bson:"favoriteProblemList"`

	JoinedClassroomList  []ClassroomID `json:"joinedClassroomList" bson:"joinedClassroomList"`
	OwnClassroomList     []ClassroomID `json:"ownClassroomList" bson:"ownClassroomList"`
	AppliedClassroomList []ClassroomID `json:"appliedClassroomList" bson:"appliedClassroomList"`
}
