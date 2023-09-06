package dao

import (
	"fmt"
	"go-web/entity"
	"time"
)

func CourseUpdate(cid int, coursename string, credit int, starttime time.Time) bool {
	mp := map[string]interface{}{
		"name":      coursename,
		"credit":    credit,
		"starttime": starttime,
	}
	db.Table("course").Where("courseid=?", cid).Updates(mp)
	return true
}
func CourseAdd(coursename string, credit int, starttime time.Time, Tno string) bool {
	cour := entity.Course{}
	db.Table("course").Last(&cour)
	db.Table("course").Create(entity.NewCourse(cour.CourseID+1, coursename, credit, starttime, Tno))
	return true
}
func CourseDelete(cno int) bool {
	//该课程的信息
	res := db.Table("course").Where("courseid=?", cno)
	if res.RowsAffected != 0 {
		return false
	}
	//删除该课程对应的作业
	assignment := []entity.Assignment{}
	db.Table("assignment").Where(" courseid=?", cno).Find(&assignment)
	for i := 0; i < len(assignment); i++ {
		studentassignment := []entity.Studentassignment{}
		db.Table("studentassignment").Where("assignmentid=?", assignment[i].AssignmentID).Delete(&studentassignment)
	}
	db.Table("assignment").Where(" courseid=?", cno).Delete(&assignment)
	//删除学生对应课表
	classschedule := []entity.ClassSchedule{}
	db.Table("classschedule").Where("courseid=?", cno).Delete(&classschedule)

	//删除该课程
	course := entity.Course{}
	db.Table("course").Where("courseid=?", cno).Delete(&course)
	return true
}

//查询该课程的作业提交情况
func CourseAssignmentQuery(cno int) []entity.AssignmentCondition {
	sche := []entity.ClassSchedule{}
	res := []entity.AssignmentCondition{}
	//查询学生课表
	db.Table("classschedule").Where("courseid=?", cno).Find(&sche)
	for i := 0; i < len(sche); i++ {
		fmt.Println("课程名为：", sche[i])
		res = append(res, AssignmentUploadCondition(sche[i].Sno, sche[i].CourseName)...)
	}
	return res
}
func CourseQuery(cno int) []entity.Assignment {
	var res []entity.Assignment
	db.Table("assignment").Where("courseid=?", cno).Find(&res)
	return res
}
func CourseInfo(cno int) []entity.Course {
	res := []entity.Course{}
	db.Table("course").Where("courseid=?", cno).First(&res)
	return res
}
func CourseSelectAll() []entity.Course {
	res := []entity.Course{}
	db.Table("course").Find(&res)
	return res
}
