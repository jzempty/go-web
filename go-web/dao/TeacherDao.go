package dao

//
import (
	"fmt"
	"go-web/entity"
)

func TeacherAdd(ts entity.Teacher) bool {
	fmt.Println("进入dao.Tadd", ts.Tno)
	var teachers entity.Teacher
	tno := ts.Tno

	result := db.Table("Teacher").Where("tno=?", tno).First(&teachers)
	if result.RowsAffected != 0 {
		//账户已经存在,插入失败
		return false
	}
	fmt.Println("maybe")
	//插入教师信息
	db.Table("Teacher").Create(ts)
	return true
}
func TeacherDelete(tno string) bool {
	tx := db.Begin()
	var an entity.TeacherScedule
	tx.Table("course").Where("tno=?", tno).Delete(an)
	var teacher []entity.Teacher
	result := tx.Table("teacher").Where("tno=? ", tno).Delete(&teacher)
	if result.Error == nil {
		tx.Commit()
		return true
	} else {
		tx.Rollback()
		//账户或密码错误，删除失败
		return false
	}

}
func TeacherUpdate(tno string, mp map[string]interface{}) bool {
	db.Table("teacher").Where("tno=?", tno).Updates(mp)
	return true

}
func TeacherLoginCheck(tno string, password string) []entity.Teacher {
	var t []entity.Teacher
	result := db.Table("teacher").Where("Tno=? AND password=?", tno, password).Find(&t)
	if result.RowsAffected != 0 {
		return t
	} else {
		//账户或密码错误
		return nil
	}
}
func TeacherQuery(tno string) []entity.Teacher {
	var t []entity.Teacher
	result := db.Table("teacher").Where("Tno=?", tno).Find(&t)
	if result.RowsAffected != 0 {
		return t
	} else {
		//账户或密码错误
		return nil
	}
}
func TeacherQueryAll() []entity.Teacher {
	var t []entity.Teacher
	result := db.Table("teacher").Find(&t)
	if result.RowsAffected != 0 {
		return t
	} else {
		//账户或密码错误
		return nil
	}
}
func TeacherCourseQuery(tno string) []entity.Course {
	var course []entity.Course
	db.Table("course").Where("Tno=?", tno).Find(&course)

	return course
}

func TeacherAddEmail(username string, password string, email string) bool {
	mp := map[string]interface{}{
		"email":    email,
		"password": password,
	}
	db.Table("teacher").Where("tno=?", username).Updates(mp)
	return true
}
