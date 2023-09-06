package entity

type Admin struct {
	AdminId  int    `gorm:"type:int;column:adminId;primary_key"`
	UserName string `gorm:"comment('')" json:"UserName"`
	PassWord string `gorm:"comment('')" json:"Password"`
}
