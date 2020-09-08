package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-web-template/resp"
	"github.com/hearecho/go-web-template/utils"
	"net/http"
	"time"
)

/**
从cookie中获取认证信息，如果不含有token标志则截至请求
 */
func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		r := resp.R{}.Ok().SetPath(context.Request.URL.Path)

		token,err := context.Cookie("token")
		if token == "" || err != nil {
			r = r.SetStatus(resp.INVALID_PARAMS)
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				r = r.SetStatus(resp.ERROR_AUTH_CHECK_TOKEN_FAIL)
			} else if time.Now().Unix() > claims.ExpiresAt {
				r = r.SetStatus(resp.ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			}
		}
		if r.Code != resp.SUCCESS {
			context.JSON(http.StatusUnauthorized,r)
			context.Abort()
			return
		}
		context.Next()
	}
}
