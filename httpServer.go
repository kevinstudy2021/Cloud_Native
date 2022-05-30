package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//设置访问url路径
	http.HandleFunc("/geek", geek)
	http.HandleFunc("/localhost/healthz", healthz)
	//监听80端口启动服务
	err := http.ListenAndServe(":80", nil)
	if err != nil{
		log.Fatal(err)
	}
}


func geek(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering geek handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "======details of the http request header:=========\n")
	fmt.Println(r.Header.Get("Host"))
	fmt.Println(r.Header.Get("User-Agent"))
	fmt.Println(r.Header.Get("Connection"))

	//获取变量中VERSION并写入到response header
	envs := os.Environ()
	for _, e := range envs{
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		} else {
			if parts[0] == "VERSION" {
				w.Header().Set(parts[0],parts[1])
			}
			//fmt.Printf("%s=%s\n",string(parts[0]), string(parts[1]))
		}
	}

	//将客户端request带的header写入response header
	for k, v := range r.Header {
		w.Header().Set(fmt.Sprintf("%s\n",k), fmt.Sprintf("%s\n", v))
	}
}

func healthz(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(300)
}
