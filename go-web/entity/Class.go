package entity

type Class struct {
	ClassId        int    `gorm:"type:int;column:ClassId;primary_key"json:"ClassId"`
	Classname      string `gorm:"comment('')" json:"ClassName"`
	Cdept          string `gorm:"comment('所在学院')" json:"Cdept"`
	PrimaryTeacher string `gorm:"comment('班主任')" json:"PrimaryTeacher"`
}
