package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-web-template/resp"
	"github.com/hearecho/go-web-template/utils"
	"net/http"
)
var limiter = utils.NewIPRateLimiter(1, 5)

func IpAcessLimiter() gin.HandlerFunc  {
	return func(context *gin.Context) {
		url := context.Request.URL.Path
		r := resp.R{}.Ok().SetPath(url)
		limiter := limiter.GetLimiter(url)
		if !limiter.Allow() {
			r = r.SetStatus(resp.ERROR_TOO_MANYREQUESTS)
			context.JSON(http.StatusTooManyRequests,r)
			context.Abort()
		}
		context.Next()
	}
}
