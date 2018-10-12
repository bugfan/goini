package goini

import (
	"log"
	"testing"
)

func TestIni(t *testing.T) {
	// 测试读取配置文件
	log.Println("开始读取文本配置---testing 1")
	LoadConfig("./conf/test.conf")
	log.Println("num:", CONFIG.GetInt64("num"))
	log.Println("float:", CONFIG.GetInt64("float"))
	log.Println("admin:", CONFIG.GetString("admin"))
	log.Println("tiger:", CONFIG.GetString("tiger"))
	log.Println("tiger_int64:", CONFIG.GetInt64("tiger"))

	log.Println("zhang_default:", CONFIG.GetString("zhang"))
	log.Println("zhang_section:", CONFIG.GetSectionString(":", "zhang"))

	// 测试写配置
	// log.Println("测试写配置：", CONFIG.Append("zxy", "sleeping", ":"))
	// log.Println(CONFIG.AppendConfig("zxy", "sleeping", ":"))

	// 测试读取环境变量
	log.Println("开始读取环境变量---testing 2")
	log.Println("读取 HOME", ENV.GetString("HOME", "nothing"))
	log.Println("读取 GOPATH", ENV.GetString("GOPATH", "nothing"))
	log.Println("读取 GOPATH", ENV.GetInt("NONE-ENV", 23)) //读取不到返回默认值
	// log.Println("读取 GOPATH", ENV.GetBool("GOPATH", "nothing"))
	// log.Println("读取 GOPATH", ENV.GetFloat("GOPATH", "nothing"))
	log.Println("开始测试环境变量相关的代码:")
	m, err := ReadFile(".env")
	log.Println("ReadFile:", err, m)

	m1 := NewMyEnv()
	log.Println("NewMyEnv 去读配置文件的变量:", err, m1.Getenv("TEST"))

	log.Println("NewMyEnv 去读配置文件的变量 GetAllenv:", err, m1.GetAllenv())

	m2 := NewMyEnv(".env")
	log.Println("NewMyEnv 去读配置文件的变量 ，带文件:", err, m2.Getenv("TEST"))

	log.Println("NewMyEnv 去读配置文件的变量 GetAllenv，带文件:", err, m2.GetAllenv())
}
