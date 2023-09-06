package service

import (
	"go-web/dao"
	"go-web/entity"
	"time"
)

func CourseAdd(coursename string, credit int, starttime time.Time, Tno string) bool {
	return dao.CourseAdd(coursename, credit, starttime, Tno)
}
func CourseAssignmentQuery(cno int) []entity.AssignmentCondition {
	//查询该课程实验提交情况
	return dao.CourseAssignmentQuery(cno)
}
func CourseDelete(cno int) bool {
	return dao.CourseDelete(cno)
}
func CourseQuery(cno int) []entity.Assignment {
	return dao.CourseQuery(cno)
}
func CourseInfo(cno int) []entity.Course {
	return dao.CourseInfo(cno)
}
func CourseUpdate(cid int, coursename string, credit int, starttime time.Time) bool {
	return dao.CourseUpdate(cid, coursename, credit, starttime)
}
func CourseSelectAll() []entity.Course {
	return dao.CourseSelectAll()
}
