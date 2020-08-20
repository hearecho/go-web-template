package upload

import (
	"fmt"
	"github.com/hearecho/go-web-template/file"
	"github.com/hearecho/go-web-template/logging"
	"github.com/hearecho/go-web-template/setting"
	"github.com/hearecho/go-web-template/utils"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

/**
上传图片
*/

func GetImageFullUrl(name string) string {
	return setting.AppSetting.Image.PrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMd5(fileName)

	return fileName + ext
}

func GetImagePath() string {
	return setting.AppSetting.Image.SavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool  {
	ext := file.GetExt(fileName)
	for _,allowExt := range setting.AppSetting.Image.AllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.Image.MaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}



