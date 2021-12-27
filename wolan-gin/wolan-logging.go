package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.NoRoute(proxyHandle)
	r.GET("/ping", pingHandle)

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
		//req.URL.Host = "47.100.105.217:18830"
		req.URL.User = c.Request.URL.User
	}}
	proxy.ModifyResponse = func(r *http.Response) error {
		all, _ := ioutil.ReadAll(r.Body)
		fmt.Println("------")
		fmt.Println(r.Request.URL)
		fmt.Println("======")
		fmt.Println(string(all))
		bs := bytes.NewBufferString(string(all))
		r.Body = ioutil.NopCloser(bs)
		return nil
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func _bulk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"took":0,"errors":false}`))
	//_, _ = w.Write([]byte(`{"took":0,"errors":false,"items":[]}`))
	fmt.Println(r.Method, r.URL)
}
