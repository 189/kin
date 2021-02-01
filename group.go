package kin

import (
	"net/http"
	"path"
)

type GroupRouter struct {
	// 保持访问 engine 的能力， 因为engine 能调度 router
	engine *Engine
	prefix string
	middleWares []HandleFunc
	parent *GroupRouter
}

// 创建一个新的 GroupRouter 实例
// 所有的 GroupRouter 共享同一个 engine 实例
func (g *GroupRouter) Group(prefix string) *GroupRouter {
	engine := g.engine
	newGroup := &GroupRouter{
		engine: engine,
		prefix: g.prefix + prefix,
		parent: g,
	}
	return newGroup
}

func (g *GroupRouter) addRouter(method string, pattern string, handle HandleFunc) {
	pattern = g.prefix + pattern
	g.engine.router.addRoute(method, pattern, handle)
}

func (g *GroupRouter) Get(pattern string, handle HandleFunc) {
	g.addRouter("get", pattern, handle)
}

func (g *GroupRouter) Post(pattern string, handle HandleFunc) {
	g.addRouter("get", pattern, handle)
}



// 往中间件数组里 push 中间件
func (g *GroupRouter) Use(middleware ...HandleFunc) {
	g.middleWares = append(g.middleWares, middleware...)
}


func (g *GroupRouter) createFileServer(virtualPath string, fs http.FileSystem) HandleFunc  {
	absolutePath := path.Join(g.prefix, virtualPath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(context *Context) {
		// 判断文件是否存在
		filepath := context.Params["filepath"]
		if _, err := fs.Open(filepath); err != nil {
			context.SetStatus(http.StatusNotFound)
		}
		fileServer.ServeHTTP(context.Writer, context.Req)
	}
}

// 静态资源托管， 将virtualPath 的地址请求，转发到 root 下文件系统服务上
func (g *GroupRouter) Static(virtualPath string, root string) {
	handler :=  g.createFileServer(virtualPath, http.Dir(root))
	pattern := path.Join(virtualPath, "/*filepath")
	g.Get(pattern, handler)
}






