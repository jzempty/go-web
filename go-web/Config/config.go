package Config

import (
	"fmt"
	"go-web/dao"
	"gopkg.in/yaml.v3"
	"os"
)

type conf struct {
	ServiceConf Service `yaml:"service"`
	MysqlConf   Mysql   `yaml:"mysql"`
}
type Service struct {
	AppMode  string `yaml:"app_mode"`
	HttpPort string `yaml:"http_port"`
}
type Mysql struct {
	Db         string `yaml:"db"`
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbUser     string `yaml:"db_user"`
	DbPassWord string `yaml:"db_pass_word"`
	DbName     string `yaml:"db_name"`
	Regular    string `yaml:"regular"`
}

var Config conf

func Init() {
	// 初始化配置文件
	yamlFile, err := os.ReadFile("./Config/config.yaml")
	/*
		    f,err:=os.Open("./Config/config.yaml")
			buf := make([]byte, 1025)
			_, err = f.Read(buf)
			defer f.Close()
	*/

	//用注释的方法读取文件用yaml解析会出错
	if err != nil {
		panic(err)
	}
	//解析yaml文件
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}

	LoadData()
}

func LoadData() {
	mysqlConf := Config.MysqlConf
	dsn := mysqlConf.DbUser + ":" + mysqlConf.DbPassWord + "@" + "tcp(" + mysqlConf.DbHost + ":" + mysqlConf.DbPort + ")/" + mysqlConf.DbName + "?charset=utf8mb4&parseTime=true"
	fmt.Println("dsn=", dsn)
	dao.Init(dsn)
}
