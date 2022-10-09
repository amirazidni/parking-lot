package main

import "parking-lot/cmd/service"

func main() {
	gw := service.GatewayServer{}
	gw.StartServer()
}
