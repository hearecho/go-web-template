package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-web-template/resp"
	"github.com/hearecho/go-web-template/utils"
	"github.com/hearecho/go-web-template/web"
	"web/models"
)

type auth struct {
	Username string `form:"username" valid:"Required; MaxSize(50)"`
	Password string `form:"password" valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var a auth
	httpCode, errCode := web.BindAndValid(c, &a)

	data := make(map[string]interface{})
	r := resp.R{}.Ok().SetData(data).SetPath(c.Request.URL.Path)
	if errCode == resp.SUCCESS {
		isExit := models.CheckAuth(a.Username, a.Password)
		if isExit {
			token, err := utils.GenerateToken(a.Username, a.Password)
			if err != nil {
				r = r.SetStatus(resp.ERROR_AUTH_TOKEN)
			} else {
				data["token"] = token
				c.SetCookie("token",token,3600,"/","localhost",false,true)
			}
		} else {
			r = r.SetStatus(resp.ERROR_AUTH)
		}
	} else {
		r = r.SetStatus(errCode)
		c.JSON(httpCode, r)
		return
	}
	c.JSON(httpCode, r)
}
