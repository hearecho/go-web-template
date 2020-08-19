package main

import (
	"fmt"
	"github.com/hearecho/go-pro/go-web/pkg/setting"
	"github.com/hearecho/go-pro/go-web/routers"
	"net/http"
)

func main() {
	//endless.DefaultReadTimeOut = setting.ReadTimeOut
	//endless.DefaultWriteTimeOut = setting.WriteTimeOut
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HttpPort)
	//
	//server := endless.NewServer(endPoint,routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	logging.Info("Actual pid is %d", syscall.Getpid())
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	logging.Error("Server err: %v", err)
	//}
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

