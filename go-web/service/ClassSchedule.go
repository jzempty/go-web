package service

import (
	"go-web/dao"
	"go-web/entity"
)

func AddSchedule(sno string, cno int) bool {
	return dao.AddSchedule(sno, cno)
}
func DeleteSchedule(sno string, cno int) []entity.ClassSchedule {
	return dao.DeleteSchedule(sno, cno)
}
func DropClassSchedule(coursename string) bool {
	return dao.DropClassSchedule(coursename)
}
func DropStudentSchedule(sno string) bool {
	return dao.DropStudentSchedule(sno)
}

func ClassScheduleQuery(sno string, coursename string) entity.ClassSchedule {
	return dao.ClassScheduleQuery(sno, coursename)
}
