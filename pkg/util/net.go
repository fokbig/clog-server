package util

import (
	"log"
	"net"
)

func ConnClose(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Println("连接关闭错误!")
		return
	}
}
