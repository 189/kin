package kin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	StatusCode int
	Params map[string]string
	handles []HandleFunc
	index int
}

func NewContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: writer,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
		handles: []HandleFunc{},
		index: -1,
	}
}

// 请求地址中获取参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 请求body中获取参数
func (c *Context) PostValue(key string) string {
	return c.Req.FormValue(key)
}

// 打印所有body 参数
func (c *Context) AllPostValue() map[string][]string {
	return c.Req.PostForm;
}


// 设置状态码
func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code);
}

// 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value);
}

// JSON 格式的响应
func (c *Context) Json(code int, response interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(response); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// HTML 响应
func (c *Context) Html(code int, html string)  {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}

// 数据响应
func (c *Context) Bytes(code int, data []byte) {
	c.SetStatus(code)
	c.Writer.Write(data);
}

// 字符串响应
func (c *Context) String(status int, format string, response...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(200);
	c.Writer.Write([]byte(fmt.Sprintf(format, response...)));
}

// 洋葱顺序 执行中间件, 直至中间件栈 执行完毕
func (c *Context) Next()  {
	total := len(c.handles)
	c.index++
	for ; c.index < total; c.index++ {
		c.handles[c.index](c)
	}
}




