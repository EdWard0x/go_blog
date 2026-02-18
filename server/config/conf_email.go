package config

// Email 邮箱配置，可以登录 QQ 邮箱，打开设置，开启第三方服务服务，详情请见 https://mail.qq.com/
type Email struct {
	Host     string `json:"host" mapstructure:"host"`         // 邮件服务器地址，例如 smtp.qq.com
	Port     int    `json:"port" mapstructure:"port"`         // 邮件服务器端口，常见的如 587 (TLS) 或 465 (SSL)
	From     string `json:"from" mapstructure:"from"`         // 发件人邮箱地址
	Nickname string `json:"nickname" mapstructure:"nickname"` // 发件人昵称，用于显示在邮件中的发件人信息
	Secret   string `json:"secret" mapstructure:"secret"`     // 发件人邮箱的密码或应用专用密码，用于身份验证
	IsSSL    bool   `json:"is_ssl" mapstructure:"is_ssl"`     // 是否使用 SSL 加密连接，true 表示使用，false 表示不使用
}
