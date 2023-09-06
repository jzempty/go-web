package service

import (
	"go-web/dao"
	"go-web/entity"
)

func StudentAssignmentAdd(Aid int, SNO string) bool {
	return dao.StudentAssignmentAdd(Aid, SNO)
}
func StudentAssignmentUpdate(SNO string, COURSENAME string, EXP string) bool {
	return dao.StudentAssignmentUpdate(SNO, COURSENAME, EXP)
}
func AssignmentUploadCondition(sno string, course string) []entity.AssignmentCondition {
	return dao.AssignmentUploadCondition(sno, course)
}
func AssignmentUploaded(sno string, course string) []entity.AssignmentCondition {
	return dao.AssignmentUploaded(sno, course)
}
func AssignmentUnloaded(sno string, course string) []entity.AssignmentCondition {
	return dao.AssignmentUnloaded(sno, course)
}
func StudentAssignmentQuery(sno string) []entity.AssignmentCondition {
	return dao.StudentAssignmentQuery(sno)
}
