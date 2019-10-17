package jaegerstart

import (
	"io/ioutil"
	"net/http"

	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func LogRequestOption() ginhttp.MWOption {
	return ginhttp.MWSpanObserver(func(span opentracing.Span, req *http.Request) {
		if req.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(req.Body)
			span.LogFields(log.Object("input", string(bodyBytes)))
		}
	})
}
