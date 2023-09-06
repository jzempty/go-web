package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/service"
	"strconv"
)

func UserAdd(c *gin.Context) {
	name := c.Query("name")
	sno := c.Query("sno")
	service.UserAdd(name, sno)
	c.JSON(200, "添加成功")
}
func SendEmail(c *gin.Context) {
	type Rec struct {
		Username    string `json:"username"`
		Newpassword string `json:"newpassword"`
		Email       string `json:"email"`
	}
	var rec Rec
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &rec)
	fmt.Println("data=", string(data), "email=", rec)
	if err != nil {
		fmt.Println("error=", err)
	}
	code, err := service.SendEmail(rec.Email)
	if err != nil {
		c.JSON(200, gin.H{
			"succssion": "false",
		})
	} else {
		c.JSON(200, gin.H{
			"succession": true,
			"code":       code,
		})
	}
}
func BindEmail(c *gin.Context) {
	type Rec struct {
		Username    string `json:"username"`
		Oldpassword string `json:"oldpassword"`
		Newpassword string `json:"newpassword"`
		Email       string `json:"email"`
		Status      string `json:"status"`
		Identity    string `json:"identity"`
	}
	var rec Rec
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &rec)
	Identity, _ := strconv.Atoi(rec.Identity)
	fmt.Println("data=", string(data), "email=", rec)
	if err != nil {
		fmt.Println("error=", err)
	}
	if service.BindEmail(rec.Username, rec.Oldpassword, rec.Newpassword, Identity, rec.Email) == nil {
		c.JSON(200, gin.H{
			"succession": "true",
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}

}
func ForgetPassword(c *gin.Context) {
	//忘记密码的处理
	type Rec struct {
		Username    string `json:"username"`
		Newpassword string `json:"newpassword"`
		Status      string `json:"status"`
		Identity    string `json:"identity"`
		Email       string `json:"email"`
	}
	var rec Rec
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &rec)
	Identity, _ := strconv.Atoi(rec.Identity)
	fmt.Println("data=", string(data), "email=", rec)
	if err != nil {
		fmt.Println("error=", err)
	}
	if rec.Status != "isok" {
		c.JSON(200, gin.H{
			"succession": "false",
		})
		return
	}
	if Identity == 1 {
		mp := map[string]interface{}{
			"password": rec.Newpassword,
			"email":    rec.Email,
		}
		if service.StudentUpdate(rec.Username, mp) {
			c.JSON(200, gin.H{
				"succession": "true",
			})
		} else {
			c.JSON(200, gin.H{
				"succession": "false",
			})
		}

	} else if Identity == 2 {
		ts := service.TeacherQuery(rec.Username)
		if ts != nil {
			mp := map[string]interface{}{
				"password": rec.Newpassword,
				"email":    rec.Email,
			}
			if service.TeacherUpdate(rec.Username, mp) {
				c.JSON(200, gin.H{
					"succession": "true",
				})
			} else {
				c.JSON(200, gin.H{
					"succession": "false",
				})
			}

		}
	}

}
