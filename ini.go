package  goini

import (
	"log"
	"github.com/go-ini/ini"
)

type IniConfig struct {
	conf *ini.File
}

func (self *IniConfig) Load(filename string) {
	conf, err := ini.Load(filename)
	if err != nil {
		log.Fatal(err)
	}
	self.conf = conf
}

func (self *IniConfig) GetString(key string) string {
	return self.GetSectionString("", key)
}

func (self *IniConfig) GetInt64(key string) int64 {
	return self.GetSectionInt64("", key)
}

func (self *IniConfig) GetSectionString(section string, key string) string {
	if self.conf == nil {
		return ""
	}
	s := self.conf.Section(section)
	return s.Key(key).String()
}

func (self *IniConfig) GetSectionInt64(section string, key string) int64 {
	if self.conf == nil {
		return 0
	}
	s := self.conf.Section(section)
	v, _ := s.Key(key).Int64()
	return v
}

var Config *IniConfig

func init() {
	Config = &IniConfig{}
}

func LoadConfig(filename string) {
	Config.Load(filename)
}
