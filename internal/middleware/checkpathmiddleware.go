package middleware

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckPathMiddleware struct {
}

func NewCheckPathMiddleware() *CheckPathMiddleware {
	return &CheckPathMiddleware{}
}

func (m *CheckPathMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("check path middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, ".log") {
			http.Error(w, "download file must is log", http.StatusBadRequest)
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
