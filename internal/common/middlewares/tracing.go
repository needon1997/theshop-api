package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		span := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path)
		defer span.Finish()
		c.Set("ctx", opentracing.ContextWithSpan(context.Background(), span))
		c.Set("span", span)
		c.Next()
		span.SetTag("http-code", c.Writer.Status())
	}
}
