package borm

import (
	"fmt"
)

type MySQLConfig struct {
	User     string `yaml:"user" json:"user" xml:"user"`
	Pwd      string `yaml:"pwd" json:"pwd" xml:"pwd"`
	Db       string `yaml:"db" json:"db" xml:"db"`
	Host     string `yaml:"host" json:"host" xml:"host"`
	Port     int    `yaml:"port" json:"port" xml:"port"`
	PoolSize int    `yaml:"pool_size" json:"pool_size" xml:"pool_size"`
	Charset  string `yaml:"charset" json:"charset" xml:"charset"`
	Loc      string `yaml:"loc" json:"loc" xml:"loc"`
}

func (mc *MySQLConfig) DriverName() string {
	return "mysql"
}

func (mc *MySQLConfig) ConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=%s", mc.User, mc.Pwd, mc.Host, mc.defPort(), mc.Db, mc.defCharset(), mc.defLoc())
}

func (mc *MySQLConfig) defPort() int {
	if mc.Port == 0 {
		mc.Port = 3306
	}
	return mc.Port
}

func (mc *MySQLConfig) defPoolSize() int {
	if mc.PoolSize == 0 {
		mc.PoolSize = 10
	}
	return mc.PoolSize
}

func (mc *MySQLConfig) defCharset() string {
	if len(mc.Charset) == 0 {
		mc.Charset = "utf8"
	}
	return mc.Charset
}

func (mc *MySQLConfig) defLoc() string {
	if len(mc.Loc) == 0 {
		mc.Loc = "Asia%2FShanghai"
	}
	return mc.Loc
}
