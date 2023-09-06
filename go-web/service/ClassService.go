package service

import (
	"go-web/dao"
	"go-web/entity"
)

func ClassAdd(class entity.Class) bool {
	if len(class.Classname) <= 0 || len(class.Classname) <= 0 || len(class.Cdept) <= 0 {
		return false
	}
	return dao.ClassAdd(class)
}
func ClassUpdate(mp map[string]interface{}) bool {
	dao.ClassUpdate(mp)
	return true
}
func ClassDelete(ClassId int) bool {
	if ClassQuery(ClassId) == nil {
		return false
	}
	return dao.ClassDelete(ClassId)
}
func ClassQuery(ClassID int) []entity.Class {
	return dao.ClassQuery(ClassID)
}
func ClassSelectAll() []entity.Class {
	return dao.ClassSelectAll()
}
