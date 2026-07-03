package main

import (
	"github.com/adminbugw/ENScan_GO/common"
	"github.com/adminbugw/ENScan_GO/runner"
)

func main() {
	var enOptions common.ENOptions
	common.Flag(&enOptions)
	enOptions.Parse()
	runner.RunEnumeration(&enOptions)
}
