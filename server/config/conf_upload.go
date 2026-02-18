package config

type Upload struct {
	Size int    `json:"size" mapstructure:"size"` // 图片上传的大小，单位 MB
	Path string `json:"path" mapstructure:"path"` // 图片上传的目录
}
