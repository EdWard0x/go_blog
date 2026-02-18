package uploads

import (
	"mime/multipart"
	"server/global"
	"server/model/appTypes"
)

// WhiteImageList 定义一个白名单映射，包含支持的图片文件类型
var WhiteImageList = map[string]struct{}{
	".jpg":  {},
	".png":  {},
	".jpeg": {},
	".ico":  {},
	".tiff": {},
	".gif":  {},
	".svg":  {},
	".webp": {},
}

type OSS interface {
	UploadImage(file *multipart.FileHeader) (string, string, error)
	DeleteImage(key string) error
}

func NewOSS() OSS {
	switch global.Config.System.OssType {
	case "cloudflare":
		return &Cloudflare{}
	case "local":
		return &Local{}
	//case "qiniu":
	//	return &Qiniu{}

	default:
		return &Local{}
	}
}

// NewOssWithStorage 是根据传入的存储类型返回相应的 OSS 实例
func NewOssWithStorage(storage appTypes.Storage) OSS {
	switch storage {
	case appTypes.Local:
		return &Local{}
	//case appTypes.Qiniu:
	//	return &Qiniu{}
	case appTypes.Cloudflare:
		return &Cloudflare{}
	default:
		return &Local{}
	}
}
