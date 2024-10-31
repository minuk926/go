package cmd

import (
	"CRUD-SERVER/config"
	"CRUD-SERVER/network"
	"fmt"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {

	var c = &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	fmt.Println("server port", c.config.Server.Port, " start... ")
	c.network.ServerStart(c.config.Server.Port)

	return c
}
