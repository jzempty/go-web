package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/service"
	"strconv"
	"time"
)

func AssignmentUpdate(c *gin.Context) {
	type Rec struct {
		Title    string
		Deadline time.Time
		Content  string
		Cno      string
	}
	//获取数据
	a := Rec{}
	data, _ := c.GetRawData()
	json.Unmarshal(data, &a)
	cno, _ := strconv.Atoi(a.Cno)
	fmt.Println("a=", a, "data", string(data))
	//检验数据
	if len(a.Title) <= 0 || cno <= 0 {
		ResponseMessage(c, "false", "输入数据有误")
		return
	}

	//处理数据

	if service.AssignmentUpdate(a.Title, a.Deadline, cno) {
		c.JSON(200, gin.H{
			"succession": "true",
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}

}
func AssignmentAdd(c *gin.Context) {
	//获取数据
	type rec struct {
		Title    string
		Deadline time.Time
		Content  string
		Cno      string
	}
	a := rec{}
	data, _ := c.GetRawData()
	json.Unmarshal(data, &a)

	Cno, _ := strconv.Atoi(a.Cno)

	//检验数据
	//fmt.Println("a=?", a, "data=", string(data))
	if len(a.Title) <= 0 || a.Deadline.Before(time.Now()) || Cno <= 0 {
		ResponseMessage(c, "false", "输入数据有误")
		return
	}

	judge := service.AssignmentAdd(a.Title, a.Deadline, a.Content, Cno)
	fmt.Println("judge=?", judge)
	if judge == true {

		c.JSON(200, gin.H{
			"succession": "true",
		})
	} else {
		fmt.Println("why")
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}
}
func AssignmentDelete(c *gin.Context) {
	//	data, _ := c.GetRawData()
	//获取数据
	type Rec struct {
		Title string
		Cno   string
	}
	var a Rec
	err := c.Bind(&a)
	fmt.Println("a=", a)
	if err != nil {
		fmt.Println("绑定数据错误")
	}
	Cno, _ := strconv.Atoi(a.Cno)

	//检验数据
	if Cno == 0 || len(a.Title) <= 0 {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}

	//处理
	if service.AssignmentDelete(a.Title, Cno) {
		c.JSON(200, gin.H{
			"succession": "true",
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}
}

func AssignmentConditionQuery(c *gin.Context) {
	//获取数据
	//查询某个实验的提交情况
	type Rec struct {
		Cno   int
		Title string
	}
	var a Rec
	err := c.Bind(&a)
	fmt.Println("a=", a)
	if err != nil {
		fmt.Println("绑定数据错误")
	}
	//检验数据
	if a.Cno <= 0 || len(a.Title) <= 0 {
		ResponseMessage(c, "false", "请求数据有误")
		return
	}
	//数据处理
	res := service.AssignmentConditionQuery(a.Cno, a.Title)
	fmt.Println("res=", res)
	c.JSON(200, gin.H{
		"succession": "true",
		"data":       res,
	})
}
func AssignmentCorrect(c *gin.Context) {
	//获取数据
	type Rec struct {
		Aid   int `json:"Aid"`
		Sno   string
		Score string `json:"Score"`
	}
	var a Rec
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &a)
	fmt.Println("a=", a, "data=", string(data))
	if err != nil {
		fmt.Println("绑定数据错误", err)
	}
	Score, _ := strconv.Atoi(a.Score)

	//检验数据
	if a.Aid <= 0 || len(a.Sno) <= 0 || Score < 0 {
		ResponseMessage(c, "false", "输入的信息有误，分数不能为负数")
	}

	//处理数据
	if service.AssignmentCorrect(a.Aid, a.Sno, Score) {
		c.JSON(200, gin.H{
			"succession": "true",
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}

}
