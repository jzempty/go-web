package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/entity"
	"go-web/service"
)

func StudentAdd(c *gin.Context) {
	st := entity.Student{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &st)
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
		return
	}
	//处理数据
	if len(st.Password) <= 0 {
		st.Password = "123456"
	}
	fmt.Println("data=", string(data), "st=", st)
	if len(st.Name) == 0 || len(st.Sno) == 0 || len(st.Password) < 5 || len(st.Classname) == 0 {
		fmt.Println("创建学生输入信息错误，请重新输入")
		ResponseMessage(c, "false", "输入信息不符合规范")
		return
	}

	//添加学生
	ans := service.StudentAdd(st.Name, st.Sno, st.Password, st.Classname)
	if ans {
		ResponseMessage(c, "true", "")
	} else {
		ResponseMessage(c, "false", "账户已存在")
		return
	}
}

//该函数尚未使用
func StudentQuery(c *gin.Context) {
	fmt.Println("以进入studentQuery函数")
	name := c.PostForm("name")
	sno := c.PostForm("sno")
	st := service.StudentCheck(name, sno)
	c.HTML(200, "student.html", st)
}

func StudentDelete(c *gin.Context) {
	st := entity.Student{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &st)
	if len(st.Sno) <= 0 || len(st.Password) <= 0 {
		ResponseMessage(c, "false", "该学生信息输入错误")
		return
	}
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
	}
	service.StudentDelete(st.Sno, st.Password)
	ResponseMessage(c, "true", "")
}

func StudentUpdate(c *gin.Context) {
	var st entity.Student
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &st)
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
	}
	//检测数据
	if len(st.Sno) <= 0 || len(st.Password) < 0 || len(st.Name) <= 0 || len(st.Classname) <= 0 {
		ResponseMessage(c, "false", "该学生信息输入错误")
		return
	}
	//处理数据
	mp := map[string]interface{}{
		"name":      st.Name,
		"password":  st.Password,
		"classname": st.Classname,
	}
	if service.StudentUpdate(st.Sno, mp) {
		fmt.Println("学生信息修改成功")
		ResponseMessage(c, "true", "")
	} else {
		fmt.Println("学生信息修改失败")
		ResponseMessage(c, "false", "该学生不存在或信息输入错误")
	}

}

func StudentSelectAll(c *gin.Context) {
	st := service.StudentSelectAll()
	c.JSON(200, gin.H{
		"succession": "true",
		"data":       st,
	})

}

func StudentAssignmentQuery(c *gin.Context) {
	type Rec struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var an Rec
	data, err := c.GetRawData()
	err = json.Unmarshal(data, &an)
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
	}
	//检测数据
	if len(an.Username) <= 0 || len(an.Password) <= 0 {
		ResponseMessage(c, "false", "数据输入有误")
		return
	}

	//处理数据
	var st []entity.Student
	st = service.StudentCheck(an.Username, an.Password)
	fmt.Println("pass", an)
	if st != nil {
		fmt.Println("成功登陆")
		res1 := service.StudentAssignmentQuery(an.Username)
		fmt.Println("res1", res1)
		c.JSON(200, gin.H{
			"data":       res1,
			"username":   st[0].Name,
			"succession": "true",
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}
}
