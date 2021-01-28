package kin


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




