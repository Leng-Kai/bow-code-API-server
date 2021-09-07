package schemas

type Notification struct {
	Message string `json:"message" bson:"message"`
}