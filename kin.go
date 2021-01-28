package kin

import (
	"fmt"
	"net/http"
)


type HandleFunc func(*Context)

type Engine struct {
	// 路由映射和访问
	router *router
	// 继承 路由组 方法和属性
	*GroupRouter
	// 存储所有路由组
	groups []*GroupRouter
}

// 构造函数
func New() *Engine {
	engine :=  &Engine{router: NewRouter()}
	engine.GroupRouter = &GroupRouter{
		engine: engine,
	}
	engine.groups = []*GroupRouter{engine.GroupRouter}
	return engine
}


//
//// 注册Get路由
//func (e *Engine) Get(pattern string, handle HandleFunc)  {
//	e.router.addRoute("get", pattern, handle)
//}
//
//
//// 注册Post路由
//func (e *Engine) Post(pattern string, handle HandleFunc)  {
//	e.router.addRoute("post", pattern, handle)
//}

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