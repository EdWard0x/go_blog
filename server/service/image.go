package service

import (
	"mime/multipart"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/model/other"
	"server/model/request"
	"server/uploads"
	"server/utils"

	"gorm.io/gorm"
)

type ImageService struct{}

func (s *ImageService) ImageUpload(file *multipart.FileHeader) (string, error) {
	oss := uploads.NewOSS()
	url, filename, err := oss.UploadImage(file)
	if err != nil {
		return "", err
	}

	return url, global.DB.Create(&database.Image{
		Name:     filename,
		URL:      url,
		Category: appTypes.Null,
		Storage:  global.Config.System.Storage(),
	}).Error
}

func (s *ImageService) ImageDelete(req request.ImageDelete) error {
	if len(req.IDs) == 0 {
		return nil
	}

	var images []database.Image
	if err := global.DB.Find(&images, req.IDs).Error; err != nil {
		return err
	}
	for _, image := range images {
		if err := global.DB.Transaction(func(tx *gorm.DB) error {
			oss := uploads.NewOssWithStorage(image.Storage)
			if err := tx.Delete(&image).Error; err != nil {
				return err
			}
			return oss.DeleteImage(image.Name)
		}); err != nil {
			return err
		}
	}
	return nil
}
func (imageService *ImageService) ImageList(info request.ImageList) (list interface{}, total int64, err error) {
	db := global.DB

	if info.Name != nil {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}

	if info.Category != nil {
		category := appTypes.ToCategory(*info.Category)
		db = db.Where("category = ?", category)
	}

	if info.Storage != nil {
		storage := appTypes.ToStorage(*info.Storage)
		db = db.Where("storage = ?", storage)
	}

	option := other.MySQLOption{
		PageInfo: info.PageInfo,
		Where:    db,
	}
	return utils.MySQLPagination(&database.Image{}, option)
}
