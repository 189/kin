package kin

import (
	"fmt"
	"time"
)

func GetLogger() HandleFunc {
	return func(context *Context) {
		start := time.Now()
		context.Next()
		fmt.Printf("%v %s %s", context.StatusCode, context.Req.RequestURI, time.Since(start))
	}
}
