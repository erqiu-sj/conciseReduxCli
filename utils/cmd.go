package utils

import (
	"flag"
	"reduxCli/types"
)

func CmdHandler() types.CmdConfig {
	config := types.CmdConfig{}
	flag.BoolVar(&config.CreateHooksFile, "h", false, "是否创建对应的hooks handler")

	flag.BoolVar(&config.CheckReducer, "cr", false, "检查 reducer 是否根据配置文件更新")

	flag.BoolVar(&config.CheckHooks, "ch", false, "检查 hooks 是否根据配置文件更新")

	flag.Parse()

	return config
}
