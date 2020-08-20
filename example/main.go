package main

import (
	"fmt"
	"github.com/hearecho/go-web-template/setting"
	"net/http"
	"web/routers"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.ServerSetting.Port),
		Handler:           router,
		ReadTimeout:       setting.ServerSetting.ReadTimeOut,
		WriteTimeout:      setting.ServerSetting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
}