package config

import (
	"strconv"
	"strings"

	"gorm.io/gorm/logger"
)

// Pgsql 数据库配置
type Pgsql struct {
	Host         string `json:"host" mapstructure:"host"`                     // 数据库服务器的地址
	Port         int    `json:"port" mapstructure:"port"`                     // 数据库服务器的端口号
	Config       string `json:"config" mapstructure:"config"`                 // 数据库连接的配置参数，如驱动、字符集等
	DBName       string `json:"db_name" mapstructure:"db_name"`               // 要连接的数据库名称
	Username     string `json:"username" mapstructure:"username"`             // 用于连接数据库的用户名
	Password     string `json:"password" mapstructure:"password"`             // 用于连接数据库的密码
	MaxIdleConns int    `json:"max_idle_conns" mapstructure:"max_idle_conns"` // 最大空闲连接数，控制连接池中的空闲连接数量
	MaxOpenConns int    `json:"max_open_conns" mapstructure:"max_open_conns"` // 最大打开连接数，限制同时打开的数据库连接数量
	LogMode      string `json:"log_mode" mapstructure:"log_mode"`             // 日志模式，例如 "info" 或 "silent"，用于控制日志输出
}

func (p Pgsql) Dsn() string {
	dsn := "host=" + p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.DBName + " port=" + strconv.Itoa(p.Port) + " " + p.Config
	//fmt.Println(dsn)
	return dsn
}
func (p Pgsql) LogLevel() logger.LogLevel {
	switch strings.ToLower(p.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
