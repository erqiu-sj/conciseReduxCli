package main

import (
	"fmt"
	"reduxCli/core"
	"reduxCli/utils"
	"time"
)

func main() {
	fmt.Println(`

 ██████╗ ██████╗ ███╗   ██╗███████╗██╗███████╗███████╗    ██████╗ ███████╗██████╗ ██╗   ██╗██╗  ██╗
██╔════╝██╔═══██╗████╗  ██║██╔════╝██║██╔════╝██╔════╝    ██╔══██╗██╔════╝██╔══██╗██║   ██║╚██╗██╔╝
██║     ██║   ██║██╔██╗ ██║███████╗██║███████╗█████╗      ██████╔╝█████╗  ██║  ██║██║   ██║ ╚███╔╝ 
██║     ██║   ██║██║╚██╗██║╚════██║██║╚════██║██╔══╝      ██╔══██╗██╔══╝  ██║  ██║██║   ██║ ██╔██╗ 
╚██████╗╚██████╔╝██║ ╚████║███████║██║███████║███████╗    ██║  ██║███████╗██████╔╝╚██████╔╝██╔╝ ██╗
 ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝╚═╝╚══════╝╚══════╝    ╚═╝  ╚═╝╚══════╝╚═════╝  ╚═════╝ ╚═╝  ╚═╝
`)
	utils.Green("compiling...")
	startTime := time.Now()
	initHandler := core.InitializationProcess{PathName: utils.ConfigurationFileName}
	initHandler.CheckPath().CheckTheConfiguration().ParseTheConfiguration().GenerateHooks(true)
	endTime := time.Since(startTime)
	utils.FillGray(endTime)
}
