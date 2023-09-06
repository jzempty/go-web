package service

import (
	"fmt"
	"go-web/dao"
	"go-web/entity"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func UserAdd(name string, sno string) {
	user := entity.User{
		Name: name,
		Sno:  sno,
	}

	dao.AddUser(user)
}
func AdminCheck(name string, password string) {
	dao.AdminCheck(name, password)
}

// 发送邮件

func SendEmail(email string) (string, error) {

	// 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	// 存到redis中
	//redis.Put(core.GetConfiguration().Redis.Modules, email, vcode, 300)
	now := time.Now()
	t := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	html := fmt.Sprintf(`<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>`, email, t, vcode)

	m := gomail.NewMessage()
	m.SetHeader("From", "FB Sample"+"<"+"1174712290@qq.com"+">")
	m.SetHeader("To", "<"+email+">")
	m.SetHeader("Subject", "[我的验证码]邮箱验证") //设置邮件主题
	m.SetBody("text/html", html)          //设置邮件正文
	// 第一个参数是host 第三个参数是发送邮箱，第四个参数 是邮箱密码
	d := gomail.NewDialer("smtp.qq.com", 587, "1174712290@qq.com", "vlmamvxkqpodghjd")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("错误：", err)
		return vcode, err
	}
	return vcode, nil
}
func GetFilePath(Aid int, Sno string) string {
	assign := dao.AssignmentInfo(Aid)
	if assign == nil {
		return "err"
	}
	course := dao.CourseInfo(assign[0].CourseID)
	if course == nil {
		return "err"
	}
	st := dao.StudentQuery(Sno)
	if st == nil {
		return "err"
	}
	teacher := dao.TeacherQuery(course[0].Tno)

	if teacher == nil {
		return "err"
	}
	dst := "./Student_Experiment/" + course[0].Name + "/" + teacher[0].Tno + "/" + Sno + "_" + st[0].Name + "_" + assign[0].Title + ".pdf"
	return dst
}
func MailIsNull(username string, identity int) bool {
	return dao.MailIsNull(username, identity)
}
func BindEmail(username string, oldpassword string, newpassword string, Identity int, email string) []rune {

	if Identity == 1 {
		if StudentCheck(username, oldpassword) != nil {
			dao.StudentAddEmail(username, newpassword, email)
			return nil
		} else {
			s := "账户或密码错误"
			return []rune(s)
		}

	} else if Identity == 2 {
		if TeacherLoginCheck(username, oldpassword) != nil {
			if dao.TeacherAddEmail(username, newpassword, email) {
				return nil
			}
		} else {
			s := "账户或密码错误"
			return []rune(s)
		}
	}
	s := "未知错误"
	return []rune(s)

}
