package service

import (
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/logging"
	"mall/serializer"
)

// ShowImgsService 商品图片详情的服务
type ShowImgsService struct {
}

// Show 商品图片
func (service *ShowImgsService) Show(id string) serializer.Response {
	var imgs []model.ProductImg
	code := e.SUCCESS

	err := model.DB.Where("product_id=?", id).Find(&imgs).Error
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
		Data: serializer.BuildImgs(imgs),
	}
}
