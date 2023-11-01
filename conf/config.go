package conf

import (
	"fmt"
	"strings"

	logging "github.com/sirupsen/logrus"
	"github.com/xilepeng/gin-mall/dao"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db     string
	DbHost string
	DbPort string
	DbUser string
	// DbPassword != DbPassWord 导致
	// [error] failed to initialize database, got error Error 1045:
	// Access denied for user 'root'@'172.17.0.1' (using password: NO)
	DbPassword string
	DbName     string

	AccessKey   string
	SerectKey   string
	Bucket      string
	QiniuServer string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	EsHost  string
	EsPort  string
	EsIndex string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	//从本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadQiniu(file)
	LoadEmail(file)
	LoadEs(file)
	LoadPhotoPath(file)

	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logging.Info(err) //日志内容
		panic(err)
	}

	// MySQL 主从复制配置
	//MySQL 读 （8）主
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	//MySQL 写 （2）从
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	dao.Database(pathRead, pathWrite)

	//esConn := "http://"+EsHost+":"+EsPort //TODO 读取ES配置
	//model.EsInit(esConn)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SerectKey = file.Section("qiniu").Key("SerectKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadEs(file *ini.File) {
	EsHost = file.Section("es").Key("EsHost").String()
	EsPort = file.Section("es").Key("EsPort").String()
	EsIndex = file.Section("es").Key("EsIndex").String()
}

func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
