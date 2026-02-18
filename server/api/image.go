package api

import (
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImageApi struct {
}

func (i *ImageApi) ImageUpload(c *gin.Context) {
	//Multipart form
	form, _ := c.MultipartForm()
	files := form.File["image"]

	for _, file := range files {
		url, err := service.ServiceGroupApp.ImageUpload(file)
		if err != nil {
			global.Log.Error("Failed to upload image:", zap.Error(err))
			response.FailWithMessage("Failed to upload image", c)
			return
		}
		response.OkWithDetailed(response.ImageUpload{
			Url:     url,
			OssType: global.Config.System.OssType,
		}, "Successfully uploaded image", c)
	}
}
func (i *ImageApi) ImageDelete(c *gin.Context) {
	var reqIds request.ImageDelete
	if err := c.ShouldBindJSON(&reqIds); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.ServiceGroupApp.ImageDelete(reqIds); err != nil {
		global.Log.Error("Failed to delete image:", zap.Error(err))
		response.FailWithMessage("Failed to delete image", c)
		return
	}
	response.OkWithMessage("Successfully deleted image", c)
}
func (i *ImageApi) ImageList(c *gin.Context) {
	var pageInfo request.ImageList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	imageList, total, err := service.ServiceGroupApp.ImageList(pageInfo)
	if err != nil {
		global.Log.Error("Failed to get image list:", zap.Error(err))
		response.FailWithMessage("Failed to get image list", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  imageList,
		Total: total,
	}, c)
}
