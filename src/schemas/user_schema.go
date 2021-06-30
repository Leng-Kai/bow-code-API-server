package schemas

type UserID = string

type User struct {
    UserID                  				`json:"id" bson:"_id,omitempty"`
    RegisterType			string      	`json:"registerType" bson:"registerType"`
    Name        			string      	`json:"name" bson:"name"`
    Avatar      			string      	`json:"avatar" bson:"avatar"`
    Super       			bool        	`json:"super" bson:"super"`
    JoinedCourseList		[]CourseID    	`json:"joinedCourseList" bson:"joinedCourseList"`
	OwnCourseList   		[]CourseID		`json:"ownCourseList" bson:"ownCourseList"`
	JoinedCoursePlanList	[]CoursePlanID	`json:"joinedCoursePlanList" bson:"joinedCoursePlanList"`
	OwnCoursePlanList		[]CoursePlanID	`json:"ownCoursePlanList" bson:"ownCoursePlanList"`
	JoinedClassroomList		[]ClassroomID	`json:"joinedClassroomList" bson:"joinedClassroomList"`
	OwnClassroomList		[]ClassroomID	`json:"ownClassroomList" bson:"ownClassroomList"`
}