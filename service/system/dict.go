package service

import (
	"errors"
	"go_blog/common/global"
	"go_blog/model/request"
	"go_blog/model/response"
	"go_blog/utils"
)

type DictService struct {
}

func (DictService) GetList(page, limit int, condition request.Dict) (err error, list []response.Dict) {
	result := global.GLOBAL_DB.Model(&request.Dict{}).Limit(limit).Offset(utils.GetPage(page, limit)).Where(condition).Find(&list)

	return result.Error, list
}

func (DictService) GetTotal(condition request.Dict) (total int64, err error) {
	result := global.GLOBAL_DB.Model(&request.Dict{}).Where(condition).Count(&total)
	return total, result.Error
}

func (DictService) AddDict(dict request.Dict) (err error) {
	if HasSameDictName(dict) {
		return errors.New(global.ERROR_EXISTED_DICT)
	}
	result := global.GLOBAL_DB.Create(&dict)
	return result.Error
}

func HasSameDictName(dict request.Dict) bool {
	var count int64
	global.GLOBAL_DB.Model(&request.Dict{}).Where("dictName = ?", dict.DictName).Count(&count)
	return count > 0
}
func (DictService) HasSameDictName(dict request.Dict) bool {
	var count int64
	global.GLOBAL_DB.Model(&request.Dict{}).Where("dict_name = ?", dict.DictName).Count(&count)
	return count > 0
}

func (DictService) UpdateDict(dict request.Dict) error {
	result := global.GLOBAL_DB.Save(&dict)
	return result.Error
}

func (DictService) DeleteDict(dict request.Dict) error {
	result := global.GLOBAL_DB.Delete(&dict)
	return result.Error
}
