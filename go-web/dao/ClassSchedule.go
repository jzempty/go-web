package dao

import (
	"fmt"
	"go-web/entity"
)

func AddSchedule(sno string, cno int) bool {
	course := entity.Course{}
	stu := entity.Student{}
	ans := db.Table("course").Where("courseid=?", cno).First(&course)
	res := db.Table("student").Where("sno=?", sno).First(&stu)
	if res.RowsAffected == 0 {
		fmt.Println("课程不存在请先创建课程")
		return true
	} else if ans.RowsAffected == 0 {
		fmt.Println("该学生信息不存在，无法导入课表")

	}
	var st []entity.ClassSchedule
	db.Table("classschedule").Where("StudentID=? AND CourseID", stu.StudentID, course.CourseID).First(&st)
	if st != nil {
		//账户已经存在,插入失败
		return false
	}
	db.Table("classschedule").Create(entity.NewClassCchedule(stu.StudentID, course.CourseID, sno, course.Name))
	return false
}
func DeleteSchedule(sno string, cno int) []entity.ClassSchedule {
	var sche []entity.ClassSchedule
	res := db.Table("classschedule").Where("sno=? AND courseid=?", sno, cno).Delete(sche)
	if res.RowsAffected != 0 {
		fmt.Println(sno, "信息已从", cno, "移除")
	}
	return sche
}

func ClassScheduleQuery(sno string, coursename string) entity.ClassSchedule {
	ans := entity.ClassSchedule{}
	res := db.Table("classschedule").Where("sno=? AND coursename=?", sno, coursename).First(&ans)
	if res.RowsAffected != 0 {
		fmt.Println(sno, "信息已从", coursename, "移除")
	}
	return ans
}

func DropClassSchedule(coursename string) bool {
	res := db.Table("classschedule").Delete("name=?", coursename)
	if res.RowsAffected != 0 {
		fmt.Println(coursename, "信息已移除")
	}
	return false

}
func DropStudentSchedule(sno string) bool {
	res := db.Table("classschedule").Delete("sno=?", sno)
	if res.RowsAffected != 0 {
		fmt.Println(sno, "信息已移除")
	}
	return false
}
