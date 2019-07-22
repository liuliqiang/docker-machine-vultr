package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/liuliqiang/docker-machine-vultr"
)

var Version string

func main() {
	plugin.RegisterDriver(vultr.NewDriver("", ""))
}
