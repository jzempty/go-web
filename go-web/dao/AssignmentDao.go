package dao

import (
	"fmt"
	"go-web/entity"
	"time"
)

func AssignmentAdd(title string, deadline time.Time, content string, courseid int) bool {
	assign := entity.Assignment{}
	db.Table("assignment").Last(&assign)
	res := entity.NewAssignment(assign.AssignmentID+1, title, deadline, courseid, content)
	db.Table("assignment").Create(res)
	return true
}
func AssignmentDelete(title string, cno int) bool {
	var as []entity.Assignment
	an := entity.Assignment{}
	st := entity.Studentassignment{}
	db.Table("assignment").Where("title=? AND courseid=?", title, cno).First(&as)
	if as == nil {
		return false
	}
	db.Table("studentassignment").Where(" assignmentid=?", an.AssignmentID).Delete(&st)
	db.Table("assignment").Where("title=? AND courseid=?", title, cno).Delete(&an)
	fmt.Println("ans=", st, "an=", an)
	return true
}
func AssignmentUpdate(title string, deadline time.Time, cno int) bool {
	db.Table("assignment").Where("title=? AND courseid=?", title, cno).Update("deadline", deadline)
	return true
}
func AssignmentQuery(cno int, title string) entity.Assignment {
	//查询该课程该实验情况
	res := entity.Assignment{}
	db.Table("assignment").Where("courseid=? AND title=?", cno, title).First(&res)
	return res
}
func AssignmentInfo(Aid int) []entity.Assignment {
	res := []entity.Assignment{}
	db.Table("assignment").Where("Assignmentid=?", Aid).Find(&res)
	return res
}
