package template

import (
	"fmt"
	"reduxCli/types"
	"reduxCli/utils"
)

// importHooksDependencies 导入hooks依赖
func importHooksDependencies(config types.ConciseReduxConfig, actionName string) string {

	dep := fmt.Sprint("import { useDispatch , useSelector } from 'react-redux';", "\n", "import { bindActionCreators } from '@zealforchange/conciseredux'",
		"\n", utils.NewImportWithReducer(config, actionName, types.NewImportWithReducerOps{IntroduceStateStatement: true}),
	)
	return dep
}
func dispatchWithAction(action string) (string, string) {
	dispatchBody := fmt.Sprint("const ", "dispatchWith", utils.CapitalizeTheFirstLetter(action), " = bindActionCreators(", "\n", action, ".getCallBackAll(),useDispatch())")
	dispatchName := fmt.Sprint("dispatchWith", utils.CapitalizeTheFirstLetter(action))
	return dispatchBody, dispatchName
}
func curStateWithReducer(action string) (string, string, string, string) {
	curStateName := fmt.Sprint("curStateWith", utils.CapitalizeTheFirstLetter(action))
	curState := fmt.Sprint("const ", curStateName, " = ",
		"useSelector((state: {", action, ":", action, utils.GenerateStateTypes, " }) => {", "\n", "return state.", action,
		"\n", "})",
	)
	curStateForReduxName := fmt.Sprint("curStateWith", utils.CapitalizeTheFirstLetter(action), "ForRedux")
	curStateForRedux := fmt.Sprint("const ", curStateForReduxName, " = ", action, ".getCurState()")
	return curStateForRedux, curState, curStateName, curStateForReduxName
}
func createFunctionBody(actionName string) string {
	curStateForRedux, curState, curStateName, curStateForReduxName := curStateWithReducer(actionName)
	dispatchBody, dispatchName := dispatchWithAction(actionName)

	funcBody := fmt.Sprint("export function use",
		utils.CapitalizeTheFirstLetter(actionName), "() {", "\n\n",
		dispatchBody, "\n\n", curStateForRedux, "\n\n", curState, "\n\n",
		"return {", dispatchName, " , ", curStateName, " , ", curStateForReduxName, "}", "\n\n",
		"}",
	)
	return funcBody
}

// CreateHooks 创建hooks
func CreateHooks(actionName string, config types.ConciseReduxConfig) string {
	hooksBody := fmt.Sprint(importHooksDependencies(config, actionName), "\n",
		createFunctionBody(actionName),
	)
	return hooksBody
}
