package service

import "github.com/hearecho/go-pro/go-web/models"

/**
tag所使用的service
*/
type TagService struct {
	ID         int
	Name       string
	State      int
	CreatedBy  string
	ModifiedBy string
	PageNum    int
	PageSize   int
}

func (t *TagService) GetAll() []models.Tag{
	tags := models.GetTags(t.PageNum,t.PageSize,t.getMaps())
	return tags
}

func (t *TagService)getMaps()map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_on"] = 0
	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}
	return maps
}
