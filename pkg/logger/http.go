package logger

import (
	"net/http"
	"net/http/httputil"
)

// HTTP returns http.RoundTripper with request/response logger.
func HTTP(rt http.RoundTripper, printf func(format string, args ...interface{})) http.RoundTripper {
	return &roundTripper{
		rt:     rt,
		printf: printf,
	}
}

type roundTripper struct {
	rt     http.RoundTripper
	printf func(format string, v ...interface{})
}

func (rt *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	b, _ := httputil.DumpRequestOut(req, true)
	if len(b) != 0 {
		rt.printf("Request:\n%s", b)
	}

	resp, err := rt.rt.RoundTrip(req)

	if resp != nil {
		b, _ = httputil.DumpResponse(resp, true)
		if len(b) != 0 {
			rt.printf("Response:\n%s", b)
		}
	}

	return resp, err
}
