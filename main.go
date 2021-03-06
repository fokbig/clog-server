package main

import (
	"CLogServer/dao"
	"CLogServer/pkg/setting"
	"CLogServer/pkg/util"
	"CLogServer/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"time"
)

func init() {
	setting.Setup()
	util.Setup()
	dao.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	initRouter := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:         endPoint,
		Handler:      initRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go listener()
	log.Printf("开始监听HTTP端口 %s", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		println(err.Error())
	}
}

func listener() {
	log.Println("开始监听TCP端口：10002")
	//Listen来创建服务端
	ln, err := net.Listen("tcp", ":10002")

	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("接收连接错误:", err)
			continue
		}

		//开启新的 去处理请求
		//每个请求开启一个goroutine
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer util.ConnClose(conn)

	buf := make([]byte, 512)
	for {
		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println("读取错误:", err)
			return
		}

		fmt.Println("box:", string(buf))
	}
}
