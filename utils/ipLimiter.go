package utils

import (
	"golang.org/x/time/rate"
	"sync"
)

/**
ip速率控制简单实现,使用rate包
*/

type IpRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit //填充速率
	b   int        //令牌个数
}

/**
生成ip速率限制其
*/
func NewIPRateLimiter(r rate.Limit, b int) *IpRateLimiter {
	i := &IpRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

/**
生成对应ip的速率限制器
 */
func (i *IpRateLimiter)AddIp(ip string) *rate.Limiter  {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r,i.b)
	i.ips[ip] = limiter

	return limiter
}

/**
获取速率限制器,如果没有将
 */
func (i *IpRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIp(ip)
	}

	i.mu.Unlock()

	return limiter
}



