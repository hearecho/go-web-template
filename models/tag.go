package models

import (
	"github.com/hearecho/go-web-template/pkg/utils"
)

type Tag struct {
	utils.Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/**
gorm所支持的回调方法：
创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind
*/


func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	utils.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	utils.DB.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExitTagByName(name string) bool {
	var tag Tag
	utils.DB.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExitTagByID(id int) bool {
	var tag Tag
	utils.DB.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}
func AddTag(name string, state int, createdBy string) bool {
	utils.DB.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}
func EditTag(id int, data interface{}) bool {
	utils.DB.Model(&Tag{}).Where("id=?", id).Updates(data)
	return true
}
func DeleteTag(id int) bool {
	utils.DB.Where("id=?", id).Delete(&Tag{})
	return true
}
