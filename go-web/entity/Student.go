package entity

type Student struct {
	StudentID int    `gorm:"type:int;column:StudentId;primary_key"`
	Name      string `gorm:"comment('')" json:"Sname"`
	Sno       string `gorm:"comment('')" json:"Sno"`
	Password  string `gorm:"comment('')" json:"PassWord"`
	Classname string `gorm:"comment('')" json:"ClassName"`
	ClassId   int    `gorm:"type:int;column:ClassId;"`
	Email     string `gorm:"column:email"json:"Email"`
}

func (s Student) GetSno() string {
	return s.Sno
}

func (s Student) GetSname() string {
	return s.Name
}

func NewStudent(Id int, n string, s string, p string, c string, cId int) Student {
	S := Student{
		StudentID: Id,
		Name:      n,
		Sno:       s,
		Password:  p,
		Classname: c,
		ClassId:   cId,
	}
	return S
}
