package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-web-template/setting"
	"github.com/hearecho/go-web-template/upload"
	"net/http"
	"web/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.AppSetting.RunMode)

	//静态文件
	r.StaticFS("/upload/images", http.Dir(upload.GetFileFullPath()))
	//r.POST("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)
	//apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	//apiv1.Use(limiter.IpAcessLimiter())
	//{
	//	apiv1.GET("/tags", v1.GetTags)
	//	apiv1.POST("/tags", v1.AddTag)
	//	apiv1.PUT("/tags/:id", v1.EditTag)
	//	apiv1.DELETE("/tags/:id", v1.DeleteTag)
	//}

	return r
}
