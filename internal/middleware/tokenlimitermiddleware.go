package middleware

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	burst   = 100
	rate    = 100
	seconds = 5
)

type TokenLimiterMiddleware struct {
}

func NewTokenLimiterMiddleware() *TokenLimiterMiddleware {
	return &TokenLimiterMiddleware{}
}

func (m *TokenLimiterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("tokenlimiter middleware")
		// Passthrough to next handler if need

		store := redis.New("127.0.0.1:6379")
		fmt.Printf("redis connect: %v \n", store.Ping())
		// New tokenLimiter
		var allowed, denied int32
		limiter := limit.NewTokenLimiter(rate, burst, store, "rate-token-limiter")
		if limiter.Allow() {
			atomic.AddInt32(&allowed, 1)
		} else {
			atomic.AddInt32(&denied, 1)
			http.Error(w, "token limiter deny", http.StatusBadRequest)
			return
		}

		fmt.Printf("allowed: %d, denied: %d, qps: %d\n", allowed, denied, (allowed+denied)/seconds)

		next(w, r)
	}
}
