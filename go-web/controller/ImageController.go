package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv18/pkg/result"
	"github.com/liuhongdi/digv18/service"
	"image"
	//"github.com/liuhongdi/digv18/service"
)

type IdController struct{}

func NewIdController() IdController {
	return IdController{}
}

//存储验证码的结构
type CaptchaResult struct {
	Id         string `json:"id"`
	Base64Blob string `json:"base_64_blob"`
	//VerifyValue string `json:"code"`
}

// 生成图形化验证码
func (a *IdController) GetOne(ctx *gin.Context) {
	fmt.Println(111)
	resultRes := result.NewResult(ctx)
	id, b64s, err := service.CaptMake()
	//fmt.Println("id--", id, "aa", b64s, "00-----", err)
	if err != nil {
		resultRes.Error(1, err.Error())
	}
	imgdata, err := base64.StdEncoding.DecodeString(b64s)
	if err != nil {
		fmt.Println("imgdata:", err)
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Println("img:", err)
	}
	ctx.JSON(200, gin.H{
		"captchaID": id,
		"cpatchImg": b64s,
		"img":       img,
	})
}

//验证captcha是否正确
/*func (a *IdController) Verify(c *gin.Context) {

	id := c.PostForm("id")
	capt := c.PostForm("capt")
	resultRes := result.NewResult(c)
	if id == "" || capt == "" {
		resultRes.Error(400, "param error")
	}

	if service.CaptVerify(id, capt) == true {
		resultRes.Success("success")
	} else {
		resultRes.Error(1, "verify failed")
	}
	return
}*/
