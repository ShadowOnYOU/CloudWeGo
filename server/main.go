package main

import (
	"log"
	server0 "server/kitex_gen/server/studentservice"
)

func main() {
	svr := server0.NewServer(new(StudentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
