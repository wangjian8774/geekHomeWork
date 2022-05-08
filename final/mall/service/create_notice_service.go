package service

import (
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/logging"
	"mall/serializer"
)

// CreateNoticeService 公告创建的服务
type CreateNoticeService struct {
	Text string `form:"text" json:"text"`
}

// Create 公告创建的服务
func (service *CreateNoticeService) Create() serializer.Response {
	notice := model.Notice{
		Text: service.Text,
	}
	code := e.SUCCESS
	err := model.DB.Create(&notice).Error
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
