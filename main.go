package main

import (
	"reduxCli/core"
	"reduxCli/utils"
)

func main() {
	initHandler := core.InitializationProcess{PathName: utils.ConfigurationFileName}
	initHandler.CheckPath().CheckTheConfiguration().ParseTheConfiguration()
}
