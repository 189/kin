package kin

import (
	"net/http"
	"strings"
)

type router struct {
	root map[string]*node
	routes map[string]HandleFunc
}

func NewRouter() *router {
	return &router{
		root: map[string]*node{},
		routes: make(map[string]HandleFunc),
	}
}

// 注册路由
func (r *router) addRoute(method string, pattern string, handle HandleFunc) {
	method = strings.ToLower(method)
	key := strings.ToLower(method + "-" + pattern)
	_, ok := r.root[method]
	if !ok {
		r.root[method] = &node{}
	}
	parts := r.parsePattern(pattern)
	r.root[method].insert(pattern, parts, 0)
	r.routes[key] = handle
}

func (r *router) parsePattern(pattern string) []string {
	parts := strings.Split(pattern, "/")
	var pieces []string
	for _, part := range parts {
		if part != "" {
			pieces = append(pieces, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return pieces
}

// 根据 req.URL 取 node 和 params
func (r *router) getRoute(path string, method string) (*node, map[string]string) {
	searchParts := r.parsePattern(path)
	root, ok := r.root[strings.ToLower(method)]
	if  !ok {
		return nil, nil
	}
	// 取 最末级 *node
	n := root.search(searchParts, 0)
	if n != nil {
		parts := r.parsePattern(n.pattern)
		params := make(map[string]string)
		// 提取 params 参数 获取参数和值的映射 /user/:uid/detail =>  {uid: "xxx" }
		for index, part := range parts {
			if strings.HasPrefix(string(part[0]), ":") {
				params[part[1:]] = searchParts[index]
			} else if part[0] == '*' {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.root[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

// 根据路由地址，找路由处理函数执行
func (r *router) Handle(context *Context) {
	n, params := r.getRoute(context.Path, context.Method)
	if n != nil {
		key := strings.ToLower(context.Method + "-" + n.pattern)
		context.Params = params
		// 将 路由处理函数，也放到 Handles 中，最后由 Next方法遍历执行
		context.handles = append(context.handles, r.routes[key])
		//r.routes[key](context)
	} else {
		http.NotFound(context.Writer, context.Req)
	}
	context.Next()
}






