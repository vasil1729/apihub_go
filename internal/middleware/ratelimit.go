package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/pkg/config"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type visitor struct {
	lastSeen time.Time
	count    int
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.RWMutex
)

// RateLimit returns a rate limiting middleware
func RateLimit(cfg *config.Config) gin.HandlerFunc {
	windowDuration := time.Duration(cfg.RateLimitWindowMS) * time.Millisecond
	maxRequests := cfg.RateLimitMaxRequests

	// Cleanup old visitors periodically
	go func() {
		for {
			time.Sleep(windowDuration)
			mu.Lock()
			for ip, v := range visitors {
				if time.Since(v.lastSeen) > windowDuration {
					delete(visitors, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		v, exists := visitors[ip]
		if !exists {
			visitors[ip] = &visitor{
				lastSeen: time.Now(),
				count:    1,
			}
			mu.Unlock()
			c.Next()
			return
		}

		// Reset count if window has passed
		if time.Since(v.lastSeen) > windowDuration {
			v.count = 1
			v.lastSeen = time.Now()
			mu.Unlock()
			c.Next()
			return
		}

		// Increment count
		v.count++
		v.lastSeen = time.Now()

		if v.count > maxRequests {
			mu.Unlock()
			response.TooManyRequests(c, "Too many requests. Please try again later.")
			c.Abort()
			return
		}

		mu.Unlock()
		c.Next()
	}
}
