package config

// Website 网站信息
type Website struct {
	Logo                 string `json:"logo" mapstructure:"logo"`
	FullLogo             string `json:"full_logo" mapstructure:"full_logo"`
	Title                string `json:"title" mapstructure:"title"`                                   // 网站标题
	Slogan               string `json:"slogan" mapstructure:"slogan"`                                 // 网站标语
	SloganEn             string `json:"slogan_en" mapstructure:"slogan_en"`                           // 英文标语
	Description          string `json:"description" mapstructure:"description"`                       // 网站描述
	Version              string `json:"version" mapstructure:"version"`                               // 网站版本
	CreatedAt            string `json:"created_at" mapstructure:"created_at"`                         // 创建时间
	IcpFiling            string `json:"icp_filing" mapstructure:"icp_filing"`                         // ICP 备案
	PublicSecurityFiling string `json:"public_security_filing" mapstructure:"public_security_filing"` // 公安备案
	BilibiliURL          string `json:"bilibili_url" mapstructure:"bilibili_url"`                     // Bilibili 链接
	GiteeURL             string `json:"gitee_url" mapstructure:"gitee_url"`                           // Gitee 链接
	GithubURL            string `json:"github_url" mapstructure:"github_url"`                         // GitHub 链接
	Name                 string `json:"name" mapstructure:"name"`                                     // 昵称
	Job                  string `json:"job" mapstructure:"job"`                                       // 职业
	Address              string `json:"address" mapstructure:"address"`                               // 地址
	Email                string `json:"email" mapstructure:"email"`                                   // 邮箱
	QQImage              string `json:"qq_image" mapstructure:"qq_image"`                             // QQ 图片链接
	WechatImage          string `json:"wechat_image" mapstructure:"wechat_image"`                     // 微信图片链接
}
