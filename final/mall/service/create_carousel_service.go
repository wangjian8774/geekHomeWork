package service

import (
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/logging"
	"mall/serializer"
)

// CreateCarouselService 轮播图创建的服务
type CreateCarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

// Create 创建商品
func (service *CreateCarouselService) Create() serializer.Response {
	carousel := model.Carousel{
		ImgPath: service.ImgPath,
	}
	code := e.SUCCESS

	err := model.DB.Create(&carousel).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarousel(carousel),
	}
}
