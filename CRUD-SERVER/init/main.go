package main

import (
	"CRUD-SERVER/init/cmd"
	"flag"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found path")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
