package main

import (
	"fmt"
	"github.com/hearecho/go-web-template/logging"
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
	err := s.ListenAndServe()
	if err != nil {
		logging.Error("Web start err")
		return
	}
}