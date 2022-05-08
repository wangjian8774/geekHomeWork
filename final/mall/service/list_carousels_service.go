package service

import (
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/logging"
	"mall/serializer"
)

// ListCarouselsService 视频列表服务
type ListCarouselsService struct {
}

// List 视频列表
func (service *ListCarouselsService) List() serializer.Response {
	carousels := []model.Carousel{}
	code := e.SUCCESS

	if err := model.DB.Find(&carousels).Error; err != nil {
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
		Data:   serializer.BuildCarousels(carousels),
	}
}
