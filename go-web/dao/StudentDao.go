package dao

import (
	"fmt"
	"go-web/entity"
)

func StudentAdd(name string, sno string, password string, className string) bool {
	var classes []entity.Class
	var st entity.Student
	result := db.Table("student").Where("sno=?", sno).First(&st)
	if result.RowsAffected != 0 {
		//账户已经存在,插入失败
		return false
	}
	//查询班级信息
	db.Table("class").Where("classname=?", className).Find(&classes)
	if classes == nil {
		fmt.Println("访问的班级信息不存在")
		return false
	}
	//插入学生信息
	db.Table("student").Last(&st)
	id := st.StudentID + 1
	//fmt.Println("id=", id)
	newstudent := entity.NewStudent(id, name, sno, password, classes[0].Classname, classes[0].ClassId)
	db.Table("student").Create(newstudent)
	return true
}

func StudentQuery(sno string) []entity.Student {
	var st []entity.Student
	result := db.Table("student").Where("sno=? ", sno).First(&st)
	if result.RowsAffected != 0 {
		return st
	} else {
		//账户或密码错误
		return nil
	}
}
func StudentCheck(sno string, password string) []entity.Student {
	var st []entity.Student
	result := db.Table("student").Where("sno=? AND password=?", sno, password).First(&st)
	if result.RowsAffected != 0 {
		return st
	} else {
		//账户或密码错误
		return nil
	}
}
func StudentSelectAll() []entity.Student {
	var st []entity.Student
	db.Table("student").Find(&st)
	return st
}

func StudentUpdate(sno string, mp map[string]interface{}) bool {

	db.Table("student").Where("sno=?", sno).Updates(mp)
	return true
}
func StudentDelete(sno string, password string) bool {
	var st []entity.Student
	var stsche []entity.Studentassignment
	var clsche []entity.ClassSchedule
	db.Table("student").Where("sno=? AND password=?", sno, password).Find(&st)
	if st == nil {
		return false
	}
	db.Table("studentassignment").Where("studentid=?", st[0].StudentID).Delete(&stsche)
	db.Table("classSchedule").Where("sno=?", sno).Delete(&clsche)
	db.Table("student").Where("sno=? AND password=?", sno, password).Delete(&st)
	fmt.Println("st=", st)
	if st != nil {
		return true
	} else {
		//账户或密码错误，删除失败
		return false
	}
}
func CourseStudentQuery(Cno int) []entity.Student {
	//查询该课程拥有的学生信息
	schedule := []entity.ClassSchedule{}
	var res []entity.Student
	db.Table("classschedule").Where("courseid=?", Cno).Find(&schedule)
	for i := 0; i < len(schedule); i++ {
		mid := entity.Student{}
		db.Table("student").Where("Studentid=?", schedule[i].StudentId).First(&mid)
		res = append(res, mid)
	}
	return res
}
func StudentAddEmail(username string, password string, email string) bool {
	mp := map[string]interface{}{
		"email":    email,
		"password": password,
	}
	db.Table("student").Where("sno=? ", username).Updates(mp)
	return true
}
