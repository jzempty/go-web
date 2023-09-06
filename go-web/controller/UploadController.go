package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/service"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func UploadTimeCheck(Aid int) bool {
	assignment := service.AssignmentInfo(Aid)
	fmt.Println("assignment=", assignment)
	if len(assignment) <= 0 {
		return false
	}
	if assignment[0].Deadline.After(time.Now()) {
		return true
	}
	return false
}

func Upload(c *gin.Context) {
	type Rec struct {
		Aid int
		Sno string
	}
	var rec Rec
	Aid := c.PostForm("Aid")
	rec.Sno = c.PostForm("Sno")
	rec.Aid, _ = strconv.Atoi(Aid)
	fmt.Println("rec=", rec)
	file, _ := c.FormFile("file")
	//上传文件到项目根目录，使用原文件名
	//c.String(http.StatusOK, fmt.Sprintf("'%s' upload!", file.Filename))
	//设置文件保存路径
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()
	reg := regexp.MustCompile(".*pdf$") //六位连续的数字
	var ty string
	if reg.MatchString(file.Filename) {
		ty = ".pdf"
	}
	judge, _ := regexp.MatchString(".*xlsx", file.Filename)
	if judge {
		ty = ".xlsx"
	}
	if ty != ".pdf" {
		c.JSON(200, gin.H{
			"succssion": "false",
			"err":       "文件不是pdf类型",
		})
	}
	fmt.Println("type:---------------------------------", ty)
	if UploadTimeCheck(rec.Aid) {
		dst := service.GetFilePath(rec.Aid, rec.Sno)
		if dst == "err" {
			c.JSON(200, gin.H{
				"succession": "false",
			})
			fmt.Println("error to upload")
			return
		}
		fmt.Println("dst=", dst)
		fmt.Println("上传路径为：", dst)
		out, _ := os.Create(dst)
		if err != nil {
			return
		}
		defer out.Close()
		_, err = io.Copy(out, src)
		if err != nil {
			c.JSON(200, gin.H{
				"succession": "false",
			})
			fmt.Println("error to upload")
		} else {
			c.JSON(200, gin.H{
				"succession": "true",
			})
			if service.StudentAssignmentAdd(rec.Aid, rec.Sno) {
				fmt.Println("该作业信息保存成功")
			}
			//c.File(dst)
		}
	} else {
		fmt.Println("实验已逾期")
		c.JSON(200, gin.H{
			"succssion": "false",
			"err":       "实验已逾期",
		})
	}
}
func ImportExcel(c *gin.Context) {
	Cno := c.PostForm("Cno")
	cno, _ := strconv.Atoi(Cno)
	fmt.Println("Cno=", cno)
	file, _ := c.FormFile("file")
	cutstring := func(s string) string {
		index := strings.LastIndex(s, " +")
		r := []rune(s)[:index]
		for i := 0; i < index; i++ {
			if r[i] == ' ' || r[i] == '-' || r[i] == ':' {
				r[i] = '_'
			}
		}
		return string(r)
	}
	t := time.Now().String()
	t = cutstring(t)
	dst := "./LoadDate/" + t + file.Filename
	fmt.Println("dst=", dst, "t=", t)
	c.SaveUploadedFile(file, dst)
	service.StudentInfoLoad(cno, dst)
	c.JSON(200, gin.H{
		"succession": "true",
	})
}
func GetFile(c *gin.Context) {
	type Rec struct {
		Aid int
		Sno string
	}
	var rec Rec
	Aid := c.Query("Aid")
	rec.Sno = c.Query("Sno")
	rec.Aid, _ = strconv.Atoi(Aid)
	fmt.Println("rec=", rec)
	dst := service.GetFilePath(rec.Aid, rec.Sno)
	fmt.Println("dst=", dst)
	if dst != "err" {
		c.File(dst)
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}

}
func GetFilePath(c *gin.Context) {
	type Rec struct {
		Aid int
		Sno string
	}
	var rec Rec
	Aid := c.Query("Aid")
	rec.Sno = c.Query("Sno")
	fmt.Println("rec=", rec)
	rec.Aid, _ = strconv.Atoi(Aid)
	dst := service.GetFilePath(rec.Aid, rec.Sno)
	if dst != "err" {
		c.JSON(200, gin.H{
			"succession": "true",
			"FilePath":   dst,
		})
	} else {
		c.JSON(200, gin.H{
			"succession": "false",
		})
	}

}
