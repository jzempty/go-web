package dao

import (
	"go-web/entity"
)

func ClassAdd(class entity.Class) bool {
	st := entity.Class{}
	db.Table("class").Last(&st)
	id := st.ClassId + 1
	class.ClassId = id
	var cs []entity.Class
	db.Table("class").Where("classname=? AND cdept=?", class.Classname, class.Cdept).Find(&cs)
	if cs != nil {
		return false
	}
	db.Table("class").Create(class)
	return true
}
func ClassUpdate(mp map[string]interface{}) {
	db.Table("class").Updates(mp)
}
func ClassQuery(ClassID int) []entity.Class {
	var res []entity.Class
	db.Table("class").Where("classid=?", ClassID).Find(&res)
	return res
}
func ClassDelete(ClassId int) bool {
	var res []entity.Class
	var st []entity.Student
	db.Table("student").Where("classid=?", ClassId).Find(&st)
	for i := 0; i < len(st); i++ {
		StudentDelete(st[i].Sno, st[i].Password)
	}
	db.Table("class").Where("classid=?", ClassId).Delete(&res)

	return true
}
func ClassSelectAll() []entity.Class {
	var res []entity.Class
	db.Table("class").Find(&res)
	return res
}
