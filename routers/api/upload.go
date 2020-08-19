package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-pro/go-web/pkg/logging"
	"github.com/hearecho/go-pro/go-web/pkg/resp"
	"github.com/hearecho/go-pro/go-web/pkg/upload"
	"net/http"
)

func UploadImage(c *gin.Context) {
	r := resp.R{}.Ok().SetPath(c.Request.URL.Path)
	data := make(map[string]string)

	file,image,err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		r = r.Error()
		c.JSON(http.StatusOK, r)
	}
	if image == nil {
		r = r.SetStatus(resp.INVALID_PARAMS)
	} else {
		//ext := path.Ext(image.Filename)
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()
		src := fullPath + imageName
		if ! upload.CheckImageExt(imageName) || ! upload.CheckImageSize(file) {
			r = r.SetStatus(resp.ERROR_UPLOAD_CHECK_IMAGE_FORMAT)
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				r = r.SetStatus(resp.ERROR_UPLOAD_CHECK_IMAGE_FAIL)
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				r = r.SetStatus(resp.ERROR_UPLOAD_SAVE_IMAGE_FAIL)
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}
	c.JSON(http.StatusOK,r.SetData(data))
}
