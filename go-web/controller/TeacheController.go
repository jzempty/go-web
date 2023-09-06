package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/entity"
	"go-web/service"
)

func ResponseMessage(c *gin.Context, success string, err string) {

	c.JSON(200, gin.H{
		"succession": success,
		"err":        err,
	})
}
func TeacherAdd(c *gin.Context) {
	ts := entity.Teacher{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &ts)
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
		return
	}
	if len(ts.Password) <= 0 {
		ts.Password = "123456"
	}
	//检查数据
	if len(ts.Tno) <= 0 || len(ts.Password) <= 0 || len(ts.Tname) <= 0 || len(ts.Password) <= 0 || len(ts.Tsex) <= 0 || len(ts.Tdept) <= 0 {
		ResponseMessage(c, "false", "数据输入出错")
		return
	}

	//处理数据
	if service.TeacherAdd(ts) {
		ResponseMessage(c, "true", "添加成功")
	} else {
		ResponseMessage(c, "false", "添加失败，服务器错误或该教师已纯在")
	}
}
func TeacherSelectAll(c *gin.Context) {
	type Res struct {
		Tno        string
		Tname      string
		CourseId   int
		CourseName string
		Tdept      string
		Tsex       string
	}
	var res []Res
	st := service.TeacherQueryAll()
	fmt.Println("st", st)
	for i := 0; i < len(st); i++ {
		course := service.TeacherCourseQuery(st[i].Tno)
		for j := 0; j < len(course); j++ {
			mid := Res{
				Tno:        st[i].Tno,
				Tname:      st[i].Tname,
				CourseId:   course[j].CourseID,
				CourseName: course[j].Name,
				Tdept:      st[i].Tdept,
				Tsex:       st[i].Tsex,
			}
			res = append(res, mid)
		}
	}
	c.JSON(200, gin.H{
		"succession": "true",
		"data1":      res,
		"data":       st,
	})
}

//暂未使用
func TeacherQuery(c *gin.Context) {
	fmt.Println("以进入TeacherQuery函数")
	tno := c.PostForm("tno")
	st := service.TeacherQuery(tno)
	c.HTML(200, "Teacher.html", st)
}

//
func TeacherDelete(c *gin.Context) {
	//导入数据
	ts := entity.Teacher{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &ts)
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
		return
	}
	//检查输入数据
	if len(ts.Tno) <= 0 {
		ResponseMessage(c, "false", "数据输入有误")
		return
	}
	//数据处理
	if service.TeacherDelete(ts.Tno) {
		ResponseMessage(c, "true", "")
	}
}

func TeacherUpdate(c *gin.Context) {
	//接收数据
	ts := entity.Teacher{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &ts)
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
	}
	if len(ts.Password) <= 0 {
		ts.Password = "123456"
	}
	//检查数据
	//处理数据
	mp := map[string]interface{}{
		"teacher_name": ts.Tname,
		"gender":       ts.Tsex,
		"teacherdept":  ts.Tdept,
		"password":     ts.Password,
	}
	if service.TeacherUpdate(ts.Tno, mp) {
		ResponseMessage(c, "true", "修改成功")
	} else {
		ResponseMessage(c, "false", "添加失败，服务器错误")
	}
}

func TeacherCourseQuery(c *gin.Context) {
	//接收数据
	var Tno string
	data, _ := c.GetRawData()
	Tno = string(data)
	fmt.Println("data:", string(data))
	//json.Unmarshal(data, &Tno)
	//检查数据
	if len(Tno) <= 0 {
		ResponseMessage(c, "false", "数据输入出错")
		return
	}
	//处理数据
	res1 := service.TeacherCourseQuery(Tno)
	fmt.Println(res1)
	c.JSON(200, gin.H{
		"data":       res1,
		"succession": "true",
	})
}
