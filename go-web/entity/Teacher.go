package entity

type Teacher struct {
	Tno      string `gorm:"not null comment('教工号') INT(11)" json:"Tno"`
	Tname    string `gorm:"type:;column:teacher_name;" json:"Tname"`
	Tsex     string `gorm:"type:;column:gender;"`
	Tdept    string `gorm:"type:;column:teacherdept;"`
	Password string `gorm:"type:;column:password;"`
	Email    string `gorm:"column:email"json:"email"`
}

func (t Teacher) Gettno() string {
	return t.Tno
}
func NewTeacher(tno string, tname string, gender string, tdept string, password string) Teacher {
	T := Teacher{
		Tno:      tno,
		Tname:    tname,
		Tsex:     gender,
		Tdept:    tdept,
		Password: password,
	}
	return T
}
