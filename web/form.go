package web

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-web-template/logging"
	"github.com/hearecho/go-web-template/resp"
	"net/http"
)

/**
绑定表单并进行验证
 */
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest,resp.INVALID_PARAMS
	}
	valid := validation.Validation{}
	check,err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, resp.ERROR
	}
	if !check {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
		return http.StatusBadRequest, resp.INVALID_PARAMS
	}
	return http.StatusOK,resp.SUCCESS
}
