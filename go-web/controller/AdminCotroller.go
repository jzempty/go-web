package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/entity"
	"go-web/service"
)

func ClassAdd(c *gin.Context) {
	//绑定数据
	class := entity.Class{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &class)
	fmt.Println("data", string(data), "class", class)
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
	}
	//检验数据
	if len(class.Classname) <= 0 || len(class.Cdept) <= 0 || len(class.PrimaryTeacher) <= 0 {
		ResponseMessage(c, "false", "输入的数据不符合要求")
	}

	//处理数据
	if service.ClassAdd(class) {
		ResponseMessage(c, "true", "")
	} else {
		ResponseMessage(c, "false", "数据有误")
	}
}
func ClassUpdate(c *gin.Context) {
	//绑定数据
	class := entity.Class{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &class)
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
	}
	//检验数据
	if len(class.Classname) <= 0 || len(class.Cdept) <= 0 || len(class.PrimaryTeacher) <= 0 {
		ResponseMessage(c, "false", "输入的数据不符合要求")
		return
	}

	//处理数据
	mp := map[string]interface{}{
		"cdept":          class.Cdept,
		"classname":      class.Classname,
		"primaryteacher": class.PrimaryTeacher,
	}
	if service.ClassUpdate(mp) {
		ResponseMessage(c, "true", "")
	} else {
		ResponseMessage(c, "false", "数据有误")
	}
}
func ClassDelete(c *gin.Context) {
	//绑定数据
	class := entity.Class{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &class)
	//检验数据
	if class.ClassId <= 0 {
		ResponseMessage(c, "false", "输入的数据不符合要求")
		return
	}
	//处理数据
	if err != nil {
		ResponseMessage(c, "false", "绑定数据出错")
	}
	if service.ClassDelete(class.ClassId) {
		ResponseMessage(c, "true", "")
	} else {
		ResponseMessage(c, "false", "数据有误")
	}
}
func ClassSelectAll(c *gin.Context) {
	res := service.ClassSelectAll()
	c.JSON(200, gin.H{
		"succession": true,
		"data":       res,
	})
}
