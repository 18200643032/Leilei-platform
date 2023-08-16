package config

import (
	"Leilei-platform/database"
	"strings"
	"os"
	"gopkg.in/ini.v1"
)

type config struct {
	BaseHost,
	BasePort,
	BaseKey,
	BaseModel,
	RunModel string
}

var Config *config

func init() {
	path, _ := os.Getwd()
	file, err := ini.Load(path + "/config/conf.ini")
	if err != nil {
		panic("配置文件读取失败")
	}
	LoadService(file)
	LoadMysql(file)
}

func LoadService(file *ini.File) {
	Config = &config{
		BaseHost: file.Section("Base").Key("Host").String(),
		BasePort: file.Section("Base").Key("Port").String(),
		RunModel: file.Section("Base").Key("Mode").String(),
	}
}

func LoadMysql(file *ini.File) {
	DbUser := file.Section("Mysql").Key("DbUser").String()
	DbPassword := file.Section("Mysql").Key("DbPassword").String()
	DBHost := file.Section("Mysql").Key("DBHost").String()
	DbPort := file.Section("Mysql").Key("DbPort").String()
	Dbname := file.Section("Mysql").Key("Dbname").String()
	pathRead := strings.Join([]string{DbUser, ":", DbPassword,
		"@tcp(", DBHost, ":", DbPort, ")/", Dbname, "?charset=utf8mb4&parseTime=True"}, "")
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword,
		"@tcp(", DBHost, ":", DbPort, ")/", Dbname, "?charset=utf8mb4&parseTime=True"}, "")
	database.InitMysql(pathRead, pathWrite)

}
