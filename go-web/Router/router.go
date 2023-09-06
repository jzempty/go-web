package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/Config"
	"go-web/controller"
	"net/http"
)

func Cors(context *gin.Context) {
	method := context.Request.Method
	// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
	//context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
	fmt.Println(context.GetHeader("Origin"))
	// 必须，设置服务器支持的所有跨域请求的方法
	context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	context.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		fmt.Println(method)
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.Next()
}
func InitRouter(router *gin.Engine) {
	router.Use(Cors)
	idController := controller.NewIdController()
	//静态文件位置
	//Router.StaticFS("/webapp", http.Dir("./webapp"))
	router.LoadHTMLGlob("webapp/static/**/*")
	router.Static("webapp", "./webapp")
	//Router.LoadHTMLFiles("webapp/static/template/index.html")
	router.GET("/user/add", controller.UserAdd)
	router.GET("/file", func(c *gin.Context) { c.HTML(200, "file.html", nil) })
	router.GET("/GetFile", controller.GetFile)
	router.GET("GetFilePath", controller.GetFilePath)
	//Router.GET("/register.html", register)
	//Router.GET("/", login)
	router.GET("/do_register", controller.StudentAdd)
	router.GET("/getImage", idController.GetOne)

	//处理POST请求
	//课程处理
	router.GET("/CourseSelectAll", controller.CourseSelectAll)
	router.POST("/SendEmail", controller.SendEmail)
	router.POST("/CourseAdd", controller.CourseAdd)                         //新建课程
	router.POST("/CourseAssignmentQuery", controller.CourseAssignmentQuery) //查询该课程的作业提交信息
	router.POST("/CourseQuery", controller.CourseQuery)                     //查询课程发布的实验信息
	router.POST("/CourseDelete", controller.CourseDelete)                   //删除该课程
	router.POST("/CourseStudentQuery", controller.CourseStudentQuery)       //查询该课程的学生信息
	router.POST("/CourseUpdate", controller.CourseUpdate)                   //更新课程信息
	router.POST("/DeleteSchedule", controller.DeleteSchedule)               //删除该课程某个学生信息
	//作业处理
	router.POST("/AssignmentAdd", controller.AssignmentAdd)
	router.POST("/AssignmentUpdate", controller.AssignmentUpdate)
	router.POST("/AssignmentDelete", controller.AssignmentDelete)                 //删除某课程作业
	router.POST("/AssignmentConditionQuery", controller.AssignmentConditionQuery) //某实验提交情况
	router.POST("/AssignmentCorrect", controller.AssignmentCorrect)               //批改作业
	//导入学生信息
	router.POST("/ImportStudentInfo", controller.ImportExcel)
	//登录上传路由
	router.POST("/upload", controller.Upload)
	router.POST("/do_login", controller.LoginCheck)
	router.POST("/ForgetPassword", controller.ForgetPassword)
	router.POST("/BindEmail", controller.BindEmail)
	//学生信息处理路由

	router.POST("/StudentAdd", controller.StudentAdd)
	router.POST("/StudentUpdate", controller.StudentUpdate)
	router.POST("/StudentDelete", controller.StudentDelete)
	router.GET("/StudentSelectAll", controller.StudentSelectAll)
	router.POST("/StudentQuery", controller.StudentQuery)
	router.POST("/StudentAssignmentQuery", controller.StudentAssignmentQuery)
	//教师信息处理路由
	router.POST("/TeacherAdd", controller.TeacherAdd)
	router.POST("/TeacherUpdate", controller.TeacherUpdate)
	router.POST("/TeacherDelete", controller.TeacherDelete)
	router.GET("/TeacherSelectAll", controller.TeacherSelectAll)
	router.POST("/TeacherQuery", controller.TeacherQuery)
	router.POST("/TeacherCourseQuery", controller.TeacherCourseQuery) //查询老师的课程情况
	//班级处理路由
	router.POST("/ClassAdd", controller.ClassAdd)
	router.POST("/ClassDelete", controller.ClassDelete)
	router.GET("/ClassSelectAll", controller.ClassSelectAll)
	router.POST("/ClassUpdate", controller.ClassUpdate)
	//Router.POST("/AssignmentCheck", controller.AssignmentQuery)
	router.Run(":" + Config.Config.ServiceConf.HttpPort)

}
