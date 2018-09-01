package goini

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"`github.com/go-ini/ini`"
)

var Config *config
var ENV *env

func init() {
	Config = &config{}
	ENV = &env{}
}

// 读取纯文本配置文件代码 支持 ‘=’ ‘ ’ 等
type config struct {
	conf *ini.File
}

func (self *config) Load(filename string) error {
	conf, err := ini.Load(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	self.conf = conf
	return nil
}

func (self *config) GetString(key string) string {
	return self.GetSectionString("", key)
}

func (self *config) GetInt64(key string) int64 {
	return self.GetSectionInt64("", key)
}

// 根据指定的分隔符读取配置
func (self *config) GetSectionString(section string, key string) string {
	if self.conf == nil {
		return ""
	}
	s := self.conf.Section(section)
	return s.Key(key).String()
}

func (self *config) GetSectionInt64(section string, key string) int64 {
	if self.conf == nil {
		return 0
	}
	s := self.conf.Section(section)
	v, _ := s.Key(key).Int64()
	return v
}

// 根据分隔符写配置文件
// func (s *config) Append(key, value, section string) error {
// 	if s.conf == nil {
// 		return errors.New("File not open!")
// 	}
// 	if section == "" {
// 		section = " "
// 	}
// 	data := fmt.Sprintf("%s %s %s", key, section, value)

// 	return s.write([]byte(data))
// }

// func (s *config) AppendTo(key, value, section string) error {
// 	data := fmt.Sprintf("%s %s %s", key, section, value)
// 	return s.conf.Append(s.conf, data)
// }

// func (s *config) write(data []byte) error {
// 	s.conf.Append(s.conf, data)
// 	return nil
// }

func LoadConfig(filename string) {
	Config.Load(filename)
}

// 读取环境变量相关代码 #所有方法入参都带一个默认值，即没有此环境变量就用默认值
type env struct {
}

func (s *env) GetString(key string, backup string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return backup
}
func (s *env) GetInt(key string, backup int64) int64 {
	v := os.Getenv(key)
	if v != "" {
		iv, _ := strconv.Atoi(v)
		return int64(iv)
	}
	return backup
}
func (s *env) GetFloat(key string, backup float64) float64 {
	v := os.Getenv(key)
	if v != "" {
		fv, _ := strconv.ParseFloat(v, 64)
		return fv
	}
	return backup
}
func (s *env) GetBool(key string, backup bool) bool {
	v := os.Getenv(key)
	if v != "" {
		iv, _ := strconv.ParseBool(v)
		return iv
	}
	return backup
}
