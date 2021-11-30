package main

import (
	"fmt"
	"net/http"
	"tractor/framework/httpHandler"
)

func main() {
	r := httpHandler.New()
	r.GET("/", func(ctx *httpHandler.Context) {
		ctx.Html(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.GET("/hello", func(c *httpHandler.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *httpHandler.Context) {
		c.Json(http.StatusOK, httpHandler.HashMap{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}