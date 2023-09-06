package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-web/entity"
	"go-web/service"
	"strconv"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Identity string `json:"identity"`
	Code     string `json:"code"`
	Codeid   string `json:"codeid"`
}

var store = base64Captcha.DefaultMemStore
var a Login
var Identity int

func LoginCheck(c *gin.Context) {
	c.Bind(&a)
	verify := store.Verify(a.Codeid, a.Code, false)
	//验证登录信息
	if len(a.Username) < 0 || len(a.Password) < 6 || len(a.Password) > 16 {
		ResponseMessage(c, "false", "账户和密码格式不正确")
		return
	}

	Identity, _ = strconv.Atoi(a.Identity)
	if verify {
		if a.Identity == "1" {
			StudentLoginCheck(c)
		} else if a.Identity == "2" {
			TeacherLoginCheck(c)
		} else if a.Identity == "0" {
			AdminLoginCheck(c)
		} else {
			ResponseMessage(c, "false", "请输入正确的身份信息")
		}
	}
	return
}
func TeacherLoginCheck(c *gin.Context) {
	var st []entity.Teacher
	st = service.TeacherLoginCheck(a.Username, a.Password)
	if st != nil {
		if service.MailIsNull(a.Username, Identity) {
			c.JSON(200, gin.H{
				"succession": "true",
				"email":      "null",
			})
			return
		}
		fmt.Println("成功登陆")
		res1 := service.TeacherCourseQuery(a.Username)
		fmt.Println(res1)
		c.JSON(200, gin.H{
			"data":       res1,
			"succession": "true",
			"username":   st[0].Tname,
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}
}

func StudentLoginCheck(c *gin.Context) {
	var st []entity.Student
	st = service.StudentCheck(a.Username, a.Password)
	if st != nil {
		if service.MailIsNull(a.Username, Identity) {
			c.JSON(200, gin.H{
				"succession": "true",
				"email":      "null",
			})
			return
		}
		fmt.Println("成功登陆")
		res1 := service.StudentAssignmentQuery(a.Username)
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

func AdminLoginCheck(c *gin.Context) {
	var st []entity.Admin
	st = service.AdminLoginCheck(a.Username, a.Password)
	if st != nil {
		fmt.Println("成功登陆")
		res1 := service.StudentAssignmentQuery(a.Username)
		fmt.Println("res1", res1)
		c.JSON(200, gin.H{
			"data":       res1,
			"username":   st[0].UserName,
			"succession": "true",
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}
}
