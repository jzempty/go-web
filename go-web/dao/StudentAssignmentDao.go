package dao

import (
	"fmt"
	"go-web/entity"
	"strings"
	"time"
)

func StudentAssignmentAdd(Aid int, sno string) bool {
	st := []entity.Student{}
	db.Table("student").Where("sno=?", sno).First(&st)
	if st == nil {
		return false
	}
	NewStudentAssignment := entity.NewStudentAssignment(st[0].StudentID, Aid, time.Now(), -1)
	db.Table("StudentAssignment").Create(NewStudentAssignment)
	return true
}

func StudentAssignmentDelte(sno string, coursename string, exp string) bool {
	if len(sno) == 0 || len(coursename) == 0 || len(exp) == 0 {
		return false
		fmt.Println("信息错误")
	}
	course := entity.Course{}
	assignment := entity.Assignment{}
	st := entity.Student{}
	db.Table("course").Where("name=?", coursename).First(&course)
	db.Table("assignment").Where("courseid=?,Title=?", course.CourseID, exp).First(&assignment)
	db.Table("student").Where("sno=?", sno).First(&st)
	result := db.Table("StudentAssignment").Delete("studentid=?,assignmentid=?", st.StudentID, assignment.AssignmentID)
	if result.RowsAffected != 0 {
		return true
	} else {
		//账户或密码错误，删除失败
		return false
	}
}

func AssignmentCorrect(assignmentid int, Sno string, score int) bool {
	a := StudentQuery(Sno)
	if a == nil {
		return false
	}
	var sa []entity.Studentassignment
	db.Table("studentassignment").Where("assignmentid=? AND studentid=?", assignmentid, a[0].StudentID).Find(&sa)
	if sa == nil {
		return false
	}
	db.Table("studentassignment").Where("assignmentid=? AND studentid=?", assignmentid, a[0].StudentID).Update("score", score)
	return true
}

//

func StudentAssignmentUpdate(sno string, coursename string, exp string) bool {
	return true
}

//查询已经提交的作业信息
func AssignmentUploaded(sno string, course string) []entity.AssignmentCondition {
	res := []entity.AssignmentCondition{}
	assignmentconditions := AssignmentUploadCondition(sno, course)
	for i := 0; i < len(assignmentconditions); i++ {
		if assignmentconditions[i].Status == true {
			res = append(res, assignmentconditions[i])
		}
	}
	return res
}

//查询未提交的作业信息
func AssignmentUnloaded(sno string, course string) []entity.AssignmentCondition {
	res := []entity.AssignmentCondition{}
	assignmentconditions := AssignmentUploadCondition(sno, course)
	for i := 0; i < len(assignmentconditions); i++ {
		if assignmentconditions[i].Status == false {
			res = append(res, assignmentconditions[i])
		}
	}
	return res
}

//查询该学生该课程作业提交信息
func AssignmentUploadCondition(sno string, coursename string) []entity.AssignmentCondition {
	//获取该学生这门课的上交情况
	var assignment []entity.Assignment
	st := entity.Student{}
	var studentassignment []entity.Studentassignment
	var classschedule entity.ClassSchedule
	unupex := []entity.AssignmentCondition{}
	db.Table("student").Where("sno=?", sno).First(&st)
	db.Table("classschedule").Where("coursename=?", coursename).First(&classschedule)
	//查询该学生已经提交的实验id
	db.Table("Studentassignment").Where("studentid=?", st.StudentID).Find(&studentassignment)
	//查询该课程需要提交的实验id
	db.Table("assignment").Where("courseid=?", classschedule.CourseId).Find(&assignment)
	for i := 0; i < len(assignment); i++ {
		mid := entity.AssignmentCondition{}
		mid.Sname = st.Name
		mid.Sno = sno
		mid.Course = coursename
		mid.Cno = assignment[i].CourseID
		mid.Title = assignment[i].Title
		mid.Deadline = assignment[i].Deadline.String()
		mid.Aid = assignment[i].AssignmentID
		cutstring := func(s string) string {
			index := strings.LastIndex(s, "+")
			r := []rune(s)[:index]
			return string(r)
		}
		mid.Deadline = cutstring(mid.Deadline)
		mid.Status = false
		unupex = append(unupex, mid)
	}

	for i := 0; i < len(unupex); i++ {
		//unupex和assignment是对应的
		for j := 0; j < len(studentassignment); j++ {
			if studentassignment[j].AssignmentID == assignment[i].AssignmentID {
				//如果该实验已提交
				unupex[i].Status = true
				unupex[i].Score = studentassignment[j].Score
				break
			}
		}
	}
	return unupex
}

//查询学生的作业提交情况
func StudentAssignmentQuery(sno string) []entity.AssignmentCondition {
	//查询该学生的作业提交信息（包括以提交和未提交）
	sche := []entity.ClassSchedule{}
	res := []entity.AssignmentCondition{}
	//查询学生课表
	db.Table("classschedule").Where("sno=?", sno).Find(&sche)
	for i := 0; i < len(sche); i++ {
		fmt.Println("课程名为：", sche[i])
		res = append(res, AssignmentUploadCondition(sno, sche[i].CourseName)...)
	}
	return res
}

//查询该课程学生作业已提交的信息
func CouresAssignmentUploaded(coursename string) []entity.AssignmentCondition {

	sche := []entity.ClassSchedule{}
	res := []entity.AssignmentCondition{}
	//查询学生课表
	db.Table("classschedule").Where("coursename=?", coursename).Find(&sche)
	for i := 0; i < len(sche); i++ {
		fmt.Println("课程名为：", sche[i])
		res = append(res, AssignmentUploaded(sche[i].Sno, sche[i].CourseName)...)
	}
	return res
}

//查询该课程学生作业未提交的信息
func CouresAssignmentUnloaded(coursename string) []entity.AssignmentCondition {
	sche := []entity.ClassSchedule{}
	res := []entity.AssignmentCondition{}
	//查询学生课表
	db.Table("classschedule").Where("coursename=?", coursename).Find(&sche)
	for i := 0; i < len(sche); i++ {
		fmt.Println("课程名为：", sche[i])
		res = append(res, AssignmentUnloaded(sche[i].Sno, sche[i].CourseName)...)
	}
	return res
}
