package entity

type TeacherScedule struct {
	Tno        string `gorm:"type:int;column:Tno;"`
	CourseId   int    `gorm:"type:int;column:courseid;"`
	CourseName string `gorm:"type:int;column:coursename;"`
}
