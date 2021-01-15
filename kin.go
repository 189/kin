package kin

import (
	"fmt"
	"net/http"
)


type HandleFunc func(*Context)

type Engine struct {
	router *router
}

// 构造函数
func New() *Engine {
	return &Engine{router: NewRouter()}
}

// 注册Get路由
func (e *Engine) Get(pattern string, handle HandleFunc)  {
	e.router.AddRoute("get", pattern, handle)
}


// 注册Post路由
func (e *Engine) Post(pattern string, handle HandleFunc)  {
	e.router.AddRoute("post", pattern, handle)
}

// 处理每一次请求
func (e * Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 生成上下文对象
	context := NewContext(w, req)
	// 根据路由规则 匹配对应的回调
	e.router.Handle(context)
}

// 启动服务
func (e *Engine) Run(port string) {
	err := http.ListenAndServe(":" + port, e)
	if err != nil {
		fmt.Println("something go wrong", err)
	}
}