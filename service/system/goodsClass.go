package service

import (
	"errors"
	"go_blog/common/global"
	"go_blog/model/request"
)

type GoodsClassService struct {
}

// isExistSameClass 存在返回true
func isExistSameClass(className string) (yn bool) {
	var _goodsClass request.GoodsClassification
	global.GLOBAL_DB.Where("typeName = ?", className).First(&_goodsClass)

	return _goodsClass.ID > 0
}

// CreateGoodsClass 添加商品分类
func (gc GoodsClassService) CreateGoodsClass(goodClass request.GoodsClassification) error {
	//不能有相同名称的类型
	if isExistSameClass(goodClass.TypeName) {
		return errors.New("存在相同名称的类型")
	}
	result := global.GLOBAL_DB.Create(&goodClass)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// UpdateGoodsClass 添加商品分类
func (gc GoodsClassService) UpdateGoodsClass(goodClass request.GoodsClassification) error {
	if isExistSameClass(goodClass.TypeName) {
		return errors.New("存在相同名称的类型")
	}
	result := global.GLOBAL_DB.Model(&goodClass).Updates(goodClass)
	return result.Error

}

// DelGoodsClass 删除商品分类
func (gc GoodsClassService) DelGoodsClass(id string) error {
	if id == "" {
		return errors.New("未传id")
	} else {
		result := global.GLOBAL_DB.Where("id = ?", id).Delete(&request.GoodsClassification{})
		return result.Error
	}
}

// GetGoodsClass 获取分类列表
func (gc GoodsClassService) GetGoodsClass() []request.GoodsClassification {
	var list []request.GoodsClassification
	global.GLOBAL_DB.Find(&list)
	return list
}
