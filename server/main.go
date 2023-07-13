package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	server0 "server/kitex_gen/server/studentservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	svr := server0.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
