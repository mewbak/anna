package main

import (
	"fmt"

	"github.com/xh3b4sd/anna/core"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/language"
	"github.com/xh3b4sd/anna/server"
)

func main() {
	fmt.Printf("%s\n", "hello, I am Anna")

	textGateway := gateway.NewGateway()

	fmt.Printf("%s\n", "booting core")
	newCoreConfig := core.DefaultCoreConfig()
	newCoreConfig.TextGateway = textGateway
	newCoreConfig.LanguageNetwork = language.NewNetwork()
	newCore := core.NewCore(newCoreConfig)
	go newCore.Boot()

	fmt.Printf("%s\n", "starting server")
	newServerConfig := server.DefaultServerConfig()
	newServerConfig.TextGateway = textGateway
	newServer := server.NewServer(newServerConfig)
	go newServer.Listen()

	for {
	}
}
