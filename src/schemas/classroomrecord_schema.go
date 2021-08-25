package schemas

type ClassroomRecordID = ID

type ScoreEntry struct {
	UserID    UserID    `json:"userID" bson:"userID"`
	ProblemID ProblemID `json:"problemID" bson:"problemID"`
	Score     int       `json:"score" bson:"score"`
}

type ClassroomRecord struct {
	ClassroomRecordID ClassroomRecordID `json:"id" bson:"_id,omitempty"` // ClassroomID
	ScoreEntryList    []ScoreEntry      `json:"scoreEntryList" bson:"scoreEntryList"`
}
