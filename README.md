# Kin
A simple web framework written in golang.

# Install

`go get git@github.com:189/kin.git`

# Usage

```
func main() {
	r := kin.New()

	// 应用中间件
	r.Use(kin.GetLogger())
	// 错误恢复
	r.Use(kin.Recovery())
	// 托管静态资源目录
	r.Static("/public", "./assets")

	// Get 路由注册
	r.Get("/", func(context *kin.Context) {
		arr := []int{1, 2,3 ,4}
		context.String(http.StatusOK, "%v", arr[10])
	})

	r.Get("/users", func(context *kin.Context) {
		context.Html(http.StatusOK, "<div>users list</div>")
	})

	r.Get("/person", func(context *kin.Context){
		context.Json(http.StatusOK, &Mock{
			Name: "lily",
			Age: 22,
		})
	})

	r.Post("/post", func(context *kin.Context) {
		data := context.PostValue("name")
		fmt.Printf("%+v", data)
		j, _ := json.Marshal(data)
		context.Json(http.StatusOK, string(j))
	})

	// 应用路由组
	userGroup := r.Group("/user")
	userGroup.Get("/:name/detail", func(context *kin.Context) {
		context.Html(http.StatusOK, fmt.Sprintf("<span>name is %s</span>", context.Params))
	})
	userGroup.Get("/:name/transaction", func(context *kin.Context) {
		context.Html(http.StatusOK, "<span>html transaction response</span>")
	})

	// 指定端口运行
	r.Run("8082")
}
```