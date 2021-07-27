package main

import (
	"CarDemo1/conf"
	"CarDemo1/routes"
	"CarDemo1/service/ws"
)

func main() {
	conf.Init()
	go ws.Manager.Start()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
