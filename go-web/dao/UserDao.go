package dao

import (
	"go-web/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

func Init(dsn string) { //初始化数据库
	// 数据库账号root 密码059219 地址localhost，这里mysql只能允许localhost和127.0.0.1访问（没有设置ip)
	//dsn := "root:059219@tcp(localhost:3305)/go_web?charset=utf8&parseTime=True&loc=Local"
	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	//设置数据库连接池的配置
	sqlDB, err := db1.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	//设置数据库的隔离级别

	//
	if err != nil {
		panic(err)
	}
	db = db1
}

func AddUser(user entity.User) {
	db.AutoMigrate(&entity.User{})
	db.Table("user").Save(&user)
}

func MailIsNull(username string, identity int) bool {
	if identity == 1 {
		var st []entity.Student
		db.Table("student").Where("sno=?", username).Find(&st)
		if len(st[0].Email) <= 0 {
			return true
		}
	} else if identity == 2 {
		var ts []entity.Teacher
		db.Table("teacher").Where("tno=?", username).Find(&ts)
		if len(ts[0].Email) <= 0 {
			return true
		}
	}
	return false
}
