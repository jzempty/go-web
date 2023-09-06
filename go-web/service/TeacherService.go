package service

import (
	"go-web/dao"
	"go-web/entity"
)

func TeacherAdd(teacher entity.Teacher) bool {
	if TeacherQuery(teacher.Tno) != nil {
		return false
	}
	return dao.TeacherAdd(teacher)
}
func TeacherLoginCheck(tno string, password string) []entity.Teacher {
	return dao.TeacherLoginCheck(tno, password)
}
func TeacherQuery(tno string) []entity.Teacher {
	return dao.TeacherQuery(tno)
}
func TeacherDelete(tno string) bool {
	if TeacherQuery(tno) == nil {
		return false
	}
	return dao.TeacherDelete(tno)
}
func TeacherUpdate(tno string, ts map[string]interface{}) bool {
	//str := fmt.Sprintf("%v", ts["tno"])
	if TeacherQuery(tno) == nil {
		return false
	}
	return dao.TeacherUpdate(tno, ts)
}
func TeacherQueryAll() []entity.Teacher {
	return dao.TeacherQueryAll()
}
func TeacherCourseQuery(tno string) []entity.Course {
	return dao.TeacherCourseQuery(tno)
}
