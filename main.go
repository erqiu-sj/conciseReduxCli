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
	cmdResult := utils.CmdHandler()
	initHandler.CheckPath().CheckTheConfiguration().ParseTheConfiguration().GenerateHooks(cmdResult.CreateHooksFile)
	//cmdResult.CheckReducer
	endTime := time.Since(startTime)
	utils.FillGray(endTime)
}
