package goini

import (
	"log"
	"testing"
)

func TestIni(t *testing.T) {
	// 测试读取配置文件
	log.Println("开始读取文本配置---testing 1")
	LoadConfig("./conf/test.conf")
	log.Println("num:", Config.GetInt64("num"))
	log.Println("float:", Config.GetInt64("float"))
	log.Println("admin:", Config.GetString("admin"))
	log.Println("tiger:", Config.GetString("tiger"))
	log.Println("tiger_int64:", Config.GetInt64("tiger"))

	log.Println("zhang_default:", Config.GetString("zhang"))
	log.Println("zhang_section:", Config.GetSectionString(":", "zhang"))

	// 测试写配置
	// log.Println("测试写配置：", Config.Append("zxy", "sleeping", ":"))
	// log.Println(Config.AppendConfig("zxy", "sleeping", ":"))

	// 测试读取环境变量
	log.Println("开始读取环境变量---testing 2")
	log.Println("读取 HOME", ENV.GetString("HOME", "nothing"))
	log.Println("读取 GOPATH", ENV.GetString("GOPATH", "nothing"))
	log.Println("读取 GOPATH", ENV.GetInt("NONE-ENV", 23)) //读取不到返回默认值
	// log.Println("读取 GOPATH", ENV.GetBool("GOPATH", "nothing"))
	// log.Println("读取 GOPATH", ENV.GetFloat("GOPATH", "nothing"))
	m, err := InitEnv(".env")
	log.Println(err, m)
}
