package service

import (
	"fmt"
	"mall/cache"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/logging"
	"mall/serializer"

	"strings"
)

// ListAcceRankingService 展示配件排行的服务
type ListAcceRankingService struct {
}

// List 获取家电排行
func (service *ListAcceRankingService) List() serializer.Response {
	var products []model.Product
	code := e.SUCCESS
	// 从redis读取点击前十的视频
	pros, _ := cache.RedisClient.ZRevRange(cache.AccessoryRank, 0, 6).Result()

	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(pros, ","))
		err := model.DB.Where("id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			logging.Info(err)
			code := e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
