# cubery
a simple web framework.
## Getting started
### Getting Cubery
```
go get github.com/dsxg666/cubery
```
### Running Cubery
```go
package main

import (
	"embed"
	"net/http"

	"github.com/dsxg666/cubery"
)

//go:embed templates/*
var htmlFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	app := cubery.Default()
	// app.Static("/static", "./static")
	app.StaticFS("/static/*filepath", staticFS)
	// app.LoadHTMLGlob("templates/**/*")
	app.LoadHTMLGlobFS("templates/**/*", htmlFS)
	app.NoRoute(func(c *cubery.Context) {
		// do some 404 deal
	})
	app.GET("/", func(c *cubery.Context) {
		c.HTML(200, "who/dsxg.html", cubery.H{
			"title": "Hello",
		})
	})
	app.GET("/hello", func(c *cubery.Context) {
		// expect /hello?name=dsxg
		c.String(http.StatusOK, "hello %s!\n", c.Query("name"))
	})
	app.POST("/login", func(c *cubery.Context) {
		c.JSON(http.StatusOK, cubery.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	v1 := app.Group("/v1")
	{
		v1.GET("/hello/:name", func(c *cubery.Context) {
			// expect /hello/dsxg
			c.String(http.StatusOK, "hello %s!\n", c.Param("name"))
		})

		v1.GET("/assets/*filepath", func(c *cubery.Context) {
			c.JSON(http.StatusOK, cubery.H{"filepath": c.Param("filepath")})
		})
	}
	app.Run(":8080")
}
```
```html
{{define "who/dsxg.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h1>Dsxg</h1>
<h1>{{.title}}</h1>
{{template "public/footer.html"}}
</body>
</html>
{{end}}
```