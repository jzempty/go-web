package entity

import "time"

type Assignment struct {
	AssignmentID int    `gorm:"type:int;column:AssignmentID;primary_key"`
	Title        string `gorm:"comment('作业标题')" json:"Title"`
	Content      string
	Deadline     time.Time `gorm:"comment('截止时间')" json:"Deadline"`
	CourseID     int       `gorm:"type:int;column:courseId;"`
}

func NewAssignment(assignmentId int, title string, deadline time.Time, courseid int, content string) Assignment {
	ans := Assignment{
		AssignmentID: assignmentId,
		Title:        title,
		Deadline:     deadline,
		CourseID:     courseid,
		Content:      content,
	}
	return ans
}
