package config

type Cloudflare struct {
	Enabled         bool   `mapstructure:"enabled"`
	BucketName      string `json:"bucketName" mapstructure:"bucketName"`
	AccountId       string `json:"accountId" mapstructure:"accountId"`
	AccessKeyId     string `json:"accessKeyId" mapstructure:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" mapstructure:"accessKeySecret"`
	Cdn             string `json:"cdn" mapstructure:"cdn"`
}
