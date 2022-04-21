package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-web-template/logging"
	"github.com/hearecho/go-web-template/resp"
	"github.com/hearecho/go-web-template/upload"

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
		imageName := upload.GetFileName(image.Filename)
		fullPath := upload.GetFileFullPath()
		savePath := upload.GetFilePath()
		src := fullPath + imageName
		if ! upload.CheckFileExt(imageName) || ! upload.CheckFileSize(file) {
			r = r.SetStatus(resp.ERROR_UPLOAD_CHECK_IMAGE_FORMAT)
		} else {
			err := upload.CheckFile(fullPath)
			if err != nil {
				logging.Warn(err)
				r = r.SetStatus(resp.ERROR_UPLOAD_CHECK_IMAGE_FAIL)
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				r = r.SetStatus(resp.ERROR_UPLOAD_SAVE_IMAGE_FAIL)
			} else {
				data["image_url"] = upload.GetFileFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}
	c.JSON(http.StatusOK,r.SetData(data))
}
