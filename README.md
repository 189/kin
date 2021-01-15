# Kin
A simple web framework written in golang.

# Intall

`go get git@github.com:189/kin.git`

# Usage

```
func main() {
    r := kin.New()
    r.Get("/", func(context *kin.Context) {
        context.String(200, "good")
    })
    r.Get("/json", func(context *kin.Context) {
        mock := &MockJson{
            Name: "lily",
            Age: 22,
        }
        data, _ := json.Marshal(mock)
        context.Json(200, string(data))
    })
    
    r.Get("/html", func(context *kin.Context){
        context.Html(200, "<div>nice html response</div>")
    })
    
    r.Post("/post", func(context *kin.Context) {
        data := context.PostValue("name")
        fmt.Printf("%+v", data)
        j, _ := json.Marshal(data)
        context.Json(200, string(j))
    })
    
    r.Run("8082")
}
```