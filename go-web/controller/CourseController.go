package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/entity"
	"go-web/service"
	"os"
	"strconv"
	"strings"
	"time"
)

func CourseAdd(c *gin.Context) {
	//获取数据
	type Rec struct {
		CourseName string    `json:"CourseName"`
		Credit     string    `json:"Credit"`
		StartTime  time.Time `json:"StartDate"`
		Tno        string
		Content    string
	}
	var course Rec
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &course)
	fmt.Println("course:", course, "data=", string(data))
	if err != nil {
		ResponseMessage(c, "false", "数据接受出错")
		return
	}
	Credit, _ := strconv.Atoi(course.Credit)
	fmt.Println("course:", course)

	//检验数据
	if len(course.CourseName) == 0 || Credit <= 0 || Credit > 10 || course.StartTime.Before(time.Now()) || len(course.Tno) <= 0 {
		ResponseMessage(c, "false", "输入的课程信息错误创建课程失败，请重新输入")
		return
	}

	//处理数据
	if service.CourseAdd(course.CourseName, Credit, course.StartTime, course.Tno) {
		dst1 := "./Student_Experiment/" + course.CourseName
		_, error := os.Stat(dst1)
		if os.IsNotExist(error) {
			err1 := os.Mkdir(dst1, 0750)
			if err1 != nil {
				fmt.Println("err1=", err1)
				ResponseMessage(c, "false", "创建"+dst1+"文件夹失败")
				return
			}
		}
		dst := "./Student_Experiment/" + course.CourseName + "/" + course.Tno
		fmt.Println("dst==", dst)
		_, error = os.Stat(dst)
		if os.IsNotExist(error) {
			err := os.Mkdir(dst, 0750)
			if err != nil {
				fmt.Println("err=", err)
				ResponseMessage(c, "false", "创建"+dst+"文件夹失败")
				return
			}
		}

		c.JSON(200, gin.H{
			"succession": "true",
		})
	} else {
		ResponseMessage(c, "false", "课程已经存在，创建课程失败")
		return
	}
}

func CourseDelete(c *gin.Context) {
	//获取数据
	type Rec struct {
		Cno int
		Tno string
	}
	res := Rec{}
	err := c.Bind(&res)
	fmt.Println("res=", res)
	if err != nil {
		ResponseMessage(c, "false", "json 数据获取错误")
		return
	}
	//检验数据
	if res.Cno <= 0 || len(res.Tno) <= 0 {
		ResponseMessage(c, "false", "json 数据获取错误")
		return
	}

	//处理数据&&删除课程文件夹
	course := service.CourseInfo(res.Cno)
	if course == nil {
		ResponseMessage(c, "false", "课程不存在")
		return
	}

	//删除数据库中该课程d信息
	service.CourseDelete(res.Cno)
	fmt.Println("course=", course)
	dst := "./Student_Experiment/" + course[0].Name + "/" + res.Tno
	if len(course[0].Name) <= 0 || len(res.Tno) <= 0 {
		ResponseMessage(c, "false", "数据库信息有误")
		return
	}

	fmt.Println("dst=", dst)
	error := os.Remove(dst)
	if error != nil {
		fmt.Println("aa", error)
		ResponseMessage(c, "false", "删除文档失败，数据库信息已删除")
		return
	}

	c.JSON(200, gin.H{
		"succession": "true",
	})
}

//查看某课程的实验提交情况
func CourseAssignmentQuery(c *gin.Context) {
	//接收数据
	var Cno int
	data, _ := c.GetRawData()
	fmt.Println("data:", string(data))
	err := json.Unmarshal(data, &Cno)
	if err != nil {
		ResponseMessage(c, "false", "json 数据获取错误")
		fmt.Println("json 数据获取错误")
		return
	}
	//检验数据
	if Cno <= 0 {
		ResponseMessage(c, "false", "输入的课程号有误")
		return
	}

	//处理数据
	res := service.CourseAssignmentQuery(Cno)
	fmt.Println("res=", res)
	c.JSON(200, gin.H{
		"succession": "true",
		"data":       res,
	})
}

//查看某课程的实验
func CourseQuery(c *gin.Context) {
	//	获取数据
	var Cno int
	data, _ := c.GetRawData()
	fmt.Println("data:", string(data))
	err := json.Unmarshal(data, &Cno)
	if err != nil {
		ResponseMessage(c, "false", "json 数据获取错误")
		fmt.Println("json 数据获取错误")
		return
	}
	//检验数据
	if Cno <= 0 {
		ResponseMessage(c, "false", "输入的课程号有误")
		return
	}
	res := service.CourseQuery(Cno)
	fmt.Println("res=", res)
	c.JSON(200, gin.H{
		"succession": "true",
		"data":       res,
	})
}

//查询该课程的学生信息
func CourseStudentQuery(c *gin.Context) {
	//接收数据
	var Cno int
	data, _ := c.GetRawData()
	fmt.Println("data:", string(data))
	err := json.Unmarshal(data, &Cno)
	if err != nil {
		ResponseMessage(c, "false", "json 数据获取错误")
		fmt.Println("json 数据获取错误")
		return
	}

	//检验数据
	if Cno <= 0 {
		ResponseMessage(c, "false", "输入的课程号有误")
		return
	}
	res := service.CourseStudentQuery(Cno)
	fmt.Println("res=", res)
	c.JSON(200, gin.H{
		"succession": "true",
		"data":       res,
	})

}

//更新课程信息
func CourseUpdate(c *gin.Context) {
	//获取数据
	data, err := c.GetRawData()
	course := entity.Course{}
	json.Unmarshal(data, &course)
	d := string(data)
	fmt.Println(d)
	var a []rune
	for i := 0; i < len(d); i++ {
		index1 := strings.LastIndex(d, "credit")
		a = []rune(d)[index1+1 : index1+2]
	}
	cred, _ := strconv.Atoi(string(a))
	fmt.Println(string(a))
	course.Credit = cred
	if err != nil {
		ResponseMessage(c, "false", "json 数据获取错误")
		fmt.Println("json 数据获取错误")
		return
	}

	//检验数据
	if len(course.Name) == 0 || course.Credit <= 0 || course.Credit > 10 || course.StartTime.Before(time.Now()) {
		fmt.Println("输入的课程信息错误创建课程失败，请重新输入")
		c.JSON(200, gin.H{
			"succession": "false",
			"err":        "输入的课程信息错误创建课程失败，请重新输入",
		})
		return
	}

	//处理数据
	if service.CourseUpdate(course.CourseID, course.Name, course.Credit, course.StartTime) {
		dst := "Student_Expriment/" + course.Name
		os.Mkdir(dst, 0750)
		c.JSON(200, gin.H{
			"succession": "true",
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
			"err":        "课程已经存在，创建课程失败",
		})
	}

}

//删除该课程的学生信息
func DeleteSchedule(c *gin.Context) {
	//查询该课程的学生信息
	type rec struct {
		Sno string
		Cno string
	}
	var a rec
	data, _ := c.GetRawData()
	fmt.Println("data:", string(data))
	err := json.Unmarshal(data, &a)
	if err != nil {
		ResponseMessage(c, "false", "json 数据获取错误")
		fmt.Println("json 数据获取错误")
		return
	}
	Cno, _ := strconv.Atoi(a.Cno)
	//检验数据
	if Cno <= 0 || len(a.Sno) <= 0 {
		ResponseMessage(c, "false", "请求数据有误")
	}

	res := service.DeleteSchedule(a.Sno, Cno)
	fmt.Println("res=", res)
	c.JSON(200, gin.H{
		"succession": "true",
		"data":       res,
	})

}

//查询所有课程信息
func CourseSelectAll(c *gin.Context) {
	res := service.CourseSelectAll()
	c.JSON(200, gin.H{
		"succession": "true",
		"data":       res,
	})
}
