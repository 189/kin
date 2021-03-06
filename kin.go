package kin

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)


type HandleFunc func(*Context)

type Engine struct {
	// 路由映射和访问
	router *router
	// 继承 路由组 方法和属性
	*GroupRouter
	// 存储所有路由组
	groups []*GroupRouter
	// 模板渲染内容处理
	funcMap template.FuncMap
	// 模板
	templates *template.Template
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

// 处理每一次请求
// 执行中间件
func (e * Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandleFunc
	// 遍历所有路由组 收集中间件
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middleWares...)
		}
	}
	// 生成上下文对象
	context := NewContext(w, req)
	context.engine = e
	context.handles = middlewares
	// 根据路由规则 匹配对应的回调
	e.router.Handle(context)
}

func (e *Engine) SetFuncMap(funcmap template.FuncMap)  {
	e.funcMap = funcmap
}

// 设置 所有可用的模板
func (e *Engine) SetView(pattern string) *template.Template {
	e.templates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
	return e.templates
}

// 启动服务
func (e *Engine) Run(port string) {
	err := http.ListenAndServe(":" + port, e)
	if err != nil {
		fmt.Println("something go wrong", err)
	}
}