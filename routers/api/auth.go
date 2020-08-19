package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-pro/go-web/models"
	"github.com/hearecho/go-pro/go-web/pkg/logging"
	"github.com/hearecho/go-pro/go-web/pkg/resp"
	"github.com/hearecho/go-pro/go-web/pkg/utils"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	r := resp.R{}.Ok().SetData(data).SetPath(c.Request.URL.Path)
	if ok {
		isExit := models.CheckAuth(username,password)
		if isExit {
			token,err := utils.GenerateToken(username,password)
			if err != nil {
				r = r.SetStatus(resp.ERROR_AUTH_TOKEN)
			} else {
				data["token"] = token
			}
		} else {
			r = r.SetStatus(resp.ERROR_AUTH)
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key,err.Message)
		}
	}
	c.JSON(200, r)
}
