package kin

import (
	"net/http"
	"strings"
)

type router struct {
	routes map[string]HandleFunc
}

func NewRouter() *router {
	return &router{
		routes: make(map[string]HandleFunc),
	}
}

// 注册路由
func (r *router) AddRoute(method string, pattern string, handle HandleFunc) {
	r.routes[strings.ToLower(method + "-" + pattern)] = handle
}

// 根据路由地址，找路由处理函数执行
func (r *router) Handle(context *Context) {
	key := strings.ToLower(context.Method + "-" + context.Path)
	if route, ok := r.routes[key]; ok  {
		route(context)
	} else {
		http.NotFound(context.Writer, context.Req)
	}
}






