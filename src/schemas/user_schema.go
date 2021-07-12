package schemas

type UserID = string

type UserInfo struct {
	Name   string `json:"name" bson:"name"`
	Avatar string `json:"avatar" bson:"avatar"`
}

type User struct {
	UserID               `json:"id" bson:"_id,omitempty"`
	RegisterType         string `json:"registerType" bson:"registerType"`
	UserInfo             `json:"userInfo" bson:"userInfo"`
	Super                bool           `json:"super" bson:"super"`
	JoinedCourseList     []CourseID     `json:"joinedCourseList" bson:"joinedCourseList"`
	OwnCourseList        []CourseID     `json:"ownCourseList" bson:"ownCourseList"`
	FavoriteCourseList   []CourseID     `json:"favoriteCourseList" bson:"favoriteCourseList"`
	JoinedCoursePlanList []CoursePlanID `json:"joinedCoursePlanList" bson:"joinedCoursePlanList"`
	OwnCoursePlanList    []CoursePlanID `json:"ownCoursePlanList" bson:"ownCoursePlanList"`
	JoinedClassroomList  []ClassroomID  `json:"joinedClassroomList" bson:"joinedClassroomList"`
	OwnClassroomList     []ClassroomID  `json:"ownClassroomList" bson:"ownClassroomList"`
}
