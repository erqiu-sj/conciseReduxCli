package template

import (
	"fmt"
	"reduxCli/utils"
	"strings"
)

//  generateActionTypes 生成action
func generateActionTypes(reducerName string, actions []string) (string, string, string) {
	actionsName := fmt.Sprint(reducerName, utils.GenerateAction)
	actionsList := fmt.Sprintln("export const ", actionsName, "= {")
	for index, ele := range actions {
		if len(actions) == index+1 {
			actionsList = fmt.Sprintln(
				fmt.Sprint(actionsList),
				fmt.Sprint(strings.ToUpper(ele), ": ", "'", ele, "'", "\n", "}"),
			)
		} else {
			actionsList = fmt.Sprintln(
				fmt.Sprint(actionsList),
				fmt.Sprint(strings.ToUpper(ele), ": ", "'", ele, "'", ","),
			)
		}
	}
	actionTypesName := fmt.Sprint(reducerName, utils.GenerateActionTypes)
	actionsList =
		fmt.Sprintln(actionsList,
			"\n",
			fmt.Sprint("export type ", actionTypesName, " =  getAllValsWithActionCollectionHelper<typeof ", actionsName, ">"),
		)
	return actionsList, actionsName, actionTypesName
}

func generateStateTypes(reducerName string) (string, string) {
	stateTypesName := fmt.Sprint(reducerName, utils.GenerateStateTypes)
	content := fmt.Sprintln("export type ", stateTypesName, " = {}")
	return content, stateTypesName
}

func generateActionPayload(reducerName string) (string, string) {
	actionTypesName := fmt.Sprint(reducerName, utils.GenerateActionPayloadTypes)
	content := fmt.Sprintln("export type ", actionTypesName, "= {}")
	return content, actionTypesName
}

type reducerPayload struct {
	stateType          string
	actionPayloadTypes string
	actionTypes        string
}

func generateReducer(reducerName string, actions []string, payload reducerPayload) string {
	generateHandler := fmt.Sprint("export const ", reducerName, " = new CreateReducer<", payload.stateType, ",", payload.actionPayloadTypes, ",", payload.actionTypes, ">({})")
	generateSetReducerKey := fmt.Sprint(".setReducerKey('", reducerName, "')")
	generateActionHandler := ""
	for _, action := range actions {
		generateActionHandler = fmt.Sprint(
			generateActionHandler,
			fmt.Sprint(" .addAction('", action, "', (state, action) => {\n    return { ...state }\n  }"),
		)
	}
	generateHandler = fmt.Sprintln(
		generateHandler,
		"\n",
		generateActionHandler,
		"\n",
		generateSetReducerKey,
	)
	return generateHandler
}
func CreateReducer(reducerName string, actions []string) string {
	content := fmt.Sprintln("import { CreateReducer , getAllValsWithActionCollectionHelper } from '@zealforchange/conciseredux'\n", "")
	actionContent, _, actionTypes := generateActionTypes(reducerName, actions)
	stateContent, stateTypes := generateStateTypes(reducerName)
	actionPayload, actionPayloadTypes := generateActionPayload(reducerName)
	reducerHandler := generateReducer(reducerName, actions, reducerPayload{
		stateType:          stateTypes,
		actionPayloadTypes: actionPayloadTypes,
		actionTypes:        actionTypes,
	})
	content = fmt.Sprintln(fmt.Sprintln(content), "\n", actionContent, "\n", stateContent, "\n", actionPayload, "\n", reducerHandler)
	return content
}
