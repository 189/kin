# Kin
A simple web framework written in golang.

# Install

`go get git@github.com:189/kin.git`

# Usage

```
func formatTime(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := kin.New()

	// 应用内置的 logger 中间件
	r.Use(kin.GetLogger())
	// 应用内置recovery中间件
	r.Use(kin.Recovery())
	// 静态资源托管 /public/xx.png 请求会映射成 ./assets/xx.png,
	r.Static("/public", "./assets")

	// 设置模板目录
	r.SetView("views/*")
	// 设置模板预处理函数，可在模板文件内部调用
	r.SetFuncMap(template.FuncMap{
		"formatTime": formatTime,
	})

	// 注册Get路由 跟路由
	r.Get("/", func(context *kin.Context) {
		context.String(http.StatusOK, "Server Wake Up")
	})
	// 响应 Html
	// curl "http://localhost:8082/users"
	r.Get("/users", func(context *kin.Context) {
		context.Html(http.StatusOK, "users.html", kin.AnyMap{
			"title": "用户详情",
			"users": [2]*kin.AnyMap{
				&kin.AnyMap{
					"name": "Tom",
					"age":22,
				},
				&kin.AnyMap{
					"name": "HanMeiMei",
					"age":23,
				},
			},
		})
	})

	// 路由组 & 路由参数 & json 响应
	// curl "http://localhost:8082/user/tom/detail"
	userGroup := r.Group("/user")
	userGroup.Get("/:name/detail", func(context *kin.Context) {
		context.Json(http.StatusOK, kin.AnyMap{
			"name": context.Params["name"],
		})
	})

	r.Run("8082")
}
```