package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"go-web/dao"
	"go-web/entity"
)

func StudentInfoLoad(cno int, dst string) {
	fmt.Println("进入学生数据录入函数:")
	f, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Println("*****")
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")
	if len(rows) < 2 {
		fmt.Println("数据过少")
	}
	template := []string{"姓名", "学号", "密码", "班级"}
	for i := 0; i < len(template); i++ {
		if template[i] != rows[0][i] {
			fmt.Println("导入信息格式有误")
			return
		}
	}

	rows = rows[1:]
	for _, row := range rows {
		name := row[0]
		sno := row[1]
		password := row[2]
		classname := row[3]
		StudentAdd(name, sno, password, classname)
		AddSchedule(sno, cno)
		fmt.Println()
	}
}
func StudentAdd(name string, sno string, password string, className string) bool {
	ans := dao.StudentAdd(name, sno, password, className)
	return ans
}
func StudentQuery(sno string) []entity.Student {
	return dao.StudentQuery(sno)
}
func StudentCheck(sno string, password string) []entity.Student {
	return dao.StudentCheck(sno, password)
}

func StudentDelete(sno string, password string) bool {
	return dao.StudentDelete(sno, password)
}
func StudentUpdate(sno string, mp map[string]interface{}) bool {
	if StudentQuery(sno) == nil {
		return false
	}
	return dao.StudentUpdate(sno, mp)
}
func StudentSelectAll() []entity.Student {
	return dao.StudentSelectAll()
}
func CourseStudentQuery(Cno int) []entity.Student {
	return dao.CourseStudentQuery(Cno)
}
func AdminLoginCheck(username string, password string) []entity.Admin {
	return dao.AdminCheck(username, password)
}
