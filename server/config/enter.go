package config

type Config struct {
	Captcha Captcha `json:"captcha" mapstructure:"captcha"`
	Email   Email   `json:"email" mapstructure:"email"`
	ES      ES      `json:"es" mapstructure:"es"`
	Gaode   Gaode   `json:"gaode" mapstructure:"gaode"`
	Jwt     Jwt     `json:"jwt" mapstructure:"jwt"`
	Mysql   Mysql   `json:"mysql" mapstructure:"mysql"`
	Pgsql   Pgsql   `json:"pgsql" mapstructure:"pgsql"`
	//Qiniu   Qiniu   `json:"qiniu" mapstructure:"qiniu"`
	CF      Cloudflare `json:"cf" mapstructure:"cloudflare"`
	QQ      QQ         `json:"qq" mapstructure:"qq"`
	Redis   Redis      `json:"redis" mapstructure:"redis"`
	System  System     `json:"system" mapstructure:"system"`
	Upload  Upload     `json:"upload" mapstructure:"upload"`
	Website Website    `json:"website" mapstructure:"website"`
	Zap     Zap        `json:"zap" mapstructure:"zap"`
}
