package logger

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"go.uber.org/zap"
)

// HTTP returns http.RoundTripper with request/response logger.
func HTTP(rt http.RoundTripper, loggerName string) http.RoundTripper {
	return &roundTripper{
		rt:         rt,
		loggerName: loggerName,
	}
}

type roundTripper struct {
	rt         http.RoundTripper
	loggerName string
}

func (rt *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rl := GetLoggerFromContext(req.Context())
	if rt.loggerName != "" {
		rl = rl.Named(rt.loggerName)
	}

	if rl.Core().Enabled(zap.DebugLevel) {
		b, _ := httputil.DumpRequestOut(req, true)
		if len(b) != 0 {
			rl.Debug(fmt.Sprintf("Sending request:\n%s.", b))
		}
	} else {
		rl.Info(fmt.Sprintf("Sending request to host=%s.", req.URL.Host))
	}

	resp, err := rt.rt.RoundTrip(req)

	if err != nil { //nolint: nestif
		rl.Error("Received error", zap.Error(err))
	} else if resp != nil {
		if rl.Core().Enabled(zap.DebugLevel) {
			b, _ := httputil.DumpResponse(resp, true)
			if len(b) != 0 {
				rl.Debug(fmt.Sprintf("Received response:\n%s", b))
			}
		} else {
			rl.Info(fmt.Sprintf("Received response: %s", resp.Status))
		}
	}
	return resp, err
}
