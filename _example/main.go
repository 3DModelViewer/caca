package main

import (
	"github.com/modelhub/cc"
	"github.com/robsix/golog"
)

const (
	ccHost = "http://shflinux03:4000"
)

func main() {
	log := golog.NewConsoleLog(0)
	ccClient := cc.NewCCClient(ccHost, log)


}
