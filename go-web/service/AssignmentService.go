package service

import (
	"go-web/dao"
	"go-web/entity"
	"time"
)

func AssignmentAdd(title string, deadline time.Time, content string, courseid int) bool {
	return dao.AssignmentAdd(title, deadline, content, courseid)
}

func AssignmentDelete(title string, cno int) bool {
	return dao.AssignmentDelete(title, cno)
}

func AssignmentQuery(cno int, title string) entity.Assignment {
	//查询该课程该作业提交情况
	return dao.AssignmentQuery(cno, title)

}

func AssignmentConditionQuery(cno int, title string) []entity.AssignmentCondition {
	//查询该课程该实验的作业提交情况
	coursecondition := CourseAssignmentQuery(cno)
	res := []entity.AssignmentCondition{}
	for i := 0; i < len(coursecondition); i++ {
		if coursecondition[i].Title == title {
			res = append(res, coursecondition[i])
		}
	}
	return res
}

func AssignmentUpdate(title string, deadline time.Time, cno int) bool {
	return dao.AssignmentUpdate(title, deadline, cno)
}

func AssignmentCorrect(assignmentid int, Sno string, score int) bool {
	return dao.AssignmentCorrect(assignmentid, Sno, score)
}
func AssignmentInfo(Aid int) []entity.Assignment {
	return dao.AssignmentInfo(Aid)
}
