package logger

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// PromHTTP is a compatibility wrapper between zap's sugared logger entry
// and Prometheus HTTP logger interface.
type PromHTTP struct {
	L *zap.SugaredLogger
}

func (p *PromHTTP) Println(args ...interface{}) { p.L.Info(args...) }

// check interfaces
var (
	_ promhttp.Logger = (*PromHTTP)(nil)
)
