package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/hearecho/go-pro/go-web/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}