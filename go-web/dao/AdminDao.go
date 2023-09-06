package dao

import "go-web/entity"

func AdminCheck(username string, password string) []entity.Admin {
	var st []entity.Admin
	result := db.Table("admin").Where("username=? AND password=?", username, password).First(&st)
	if result.RowsAffected != 0 {
		return st
	} else {
		//账户或密码错误
		return nil
	}
}
