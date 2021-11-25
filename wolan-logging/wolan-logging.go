package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.NoRoute(proxyHandle)
	r.GET("/ping", pingHandle)
	r.GET("/login", proxyHandle)

	err := r.Run(":20100")
	if err != nil {
		return
	}
}

func pingHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func proxyHandle(c *gin.Context) {
	proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "svc-es:9200"
	}}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func _bulk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"took":0,"errors":false}`))
	//_, _ = w.Write([]byte(`{"took":0,"errors":false,"items":[]}`))
	fmt.Println(r.Method, r.URL)
}
