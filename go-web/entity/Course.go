package entity

import "time"

type Course struct {
	CourseID  int       `gorm:"type:int;column:CourseID;primary_key" json:"Cno"`
	Name      string    `gorm:"comment('课程名')" json:"CourseName"`
	Credit    int       `gorm:"comment('课程学分')" json:"Credit"`
	StartTime time.Time `gorm:"type:int;column:starttime;primary_key" json:"StartDate"`
	Tno       string    `gorm:"type:int;column:Tno;primary_key" json:"Tno"`
}

func NewCourse(courseid int, coursename string, credit int, starttime time.Time, Tno string) Course {
	return Course{
		CourseID:  courseid,
		Name:      coursename,
		Credit:    credit,
		StartTime: starttime,
		Tno:       Tno,
	}
}
