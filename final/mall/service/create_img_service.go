package service

import (
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/logging"
	"mall/serializer"
)

// CreateImgService 商品图片创建的服务
type CreateImgService struct {
	ProductID uint   `form:"product_id" json:"product_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

// Create 创建商品图片
func (service *CreateImgService) Create() serializer.Response {
	img := model.ProductImg{
		ProductID: service.ProductID,
		ImgPath:   service.ImgPath,
	}
	code := e.SUCCESS
	err := model.DB.Create(&img).Error
	if err != nil {
		logging.Info(err)
		code := e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
