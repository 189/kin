package kin

import "fmt"

type router struct {
	routes map[string]HandleFunc
}

func NewRouter() *router {
	return &router{
		routes: make(map[string]HandleFunc),
	}
}


func (r *router) AddRoute(method string, pattern string, handle HandleFunc) {
	r.routes[method + "-" + pattern] = handle;
}

// 根据路由地址，找路由处理函数
func (r *router) Handle(context *Context) {
	key := context.Method + "-" + context.Path;
	if route, ok := r.routes[key]; ok  {
		route(context);
	} else {
		fmt.Println("no match route")
	}
}






