package entity

import "time"

type Studentassignment struct {
	StudentID    int       `gorm:"type:int;column:studentID;"`
	AssignmentID int       `gorm:"type:int;column:AssignmentID;" json:"AssignmentID"`
	SubmitTime   time.Time `gorm:"type:int;column:submittime;" json:"SubmitTime"`
	Score        int       `gorm:"type:int;column:score;" json:"statement"`
}

func NewStudentAssignment(st int, as int, submitTime time.Time, score int) Studentassignment {
	ans := Studentassignment{
		StudentID:    st,
		AssignmentID: as,
		SubmitTime:   submitTime,
		Score:        score,
	}
	return ans
}
