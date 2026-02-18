package config

// Zap 日志配置，详情可参考七米的博客 https://liwenzhou.com/posts/Go/zap/
type Zap struct {
	Level          string `json:"level" mapstructure:"level"`                       // 日志等级，无特殊需求，用 info 即可
	Filename       string `json:"filename" mapstructure:"filename"`                 // 日志文件的位置
	MaxSize        int    `json:"max_size" mapstructure:"max_size"`                 // 在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups     int    `json:"max_backups" mapstructure:"max_backups"`           // 保留旧文件的最大个数
	MaxAge         int    `json:"max_age" mapstructure:"max_age"`                   // 保留旧文件的最大天数
	IsConsolePrint bool   `json:"is_console_print" mapstructure:"is_console_print"` // 是否在控制台打印日志，true 表示打印，false 表示不打印
}
