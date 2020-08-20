package go_web_template

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-web-template/setting"
	"net/http"
)

func Run(router *gin.Engine)  {
	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.ServerSetting.Port),
		Handler:           router,
		ReadTimeout:       setting.ServerSetting.ReadTimeOut,
		WriteTimeout:      setting.ServerSetting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
}