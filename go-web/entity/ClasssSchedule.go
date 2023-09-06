package entity

type ClassSchedule struct {
	StudentId  int    `gorm:"type:int;column:studentid;"`
	CourseId   int    `gorm:"type:int;column:courseid;"`
	Sno        string `gorm:"type:int;column:Sno;"`
	CourseName string `gorm:"type:int;column:Coursename;"`
}

func NewClassCchedule(StudentId int, CourseId int, Sno string, CourseName string) ClassSchedule {
	return ClassSchedule{
		StudentId:  StudentId,
		CourseId:   CourseId,
		Sno:        Sno,
		CourseName: CourseName,
	}
}
