package main

import (
	"reduxCli/core"
	"reduxCli/utils"
)

func main() {
	initHandler := core.InitializationProcess{PathName: utils.ConfigurationFileName}
	initHandler.CheckPath().CheckTheConfiguration().ParseTheConfiguration()
	//content := "import {applyMiddleware, combineReducers, createStore, Middleware} from '@zealforchange/conciseredux';" +
	//	"const store = createStore(\n   " +
	//	" combineReducers({\n      " +
	//	"  HeaderButtonReducer: HeaderButtonReducer.finish(),\n" +
	//	"        HeaderGoodsReducer: HeaderGoodsReducer.finish(),\n        " +
	//	"HeaderCustomer: HeaderCustomer.finish(),\n        " +
	//	"YarnSelectionListStatus: YarnSelectionListStatus.finish(),\n        " +
	//	"Common: Common.finish(),\n        " +
	//	"FormOptionsInit: FormOptionsInit.finish(),\n        LikeAndDeleteIdCollection: LikeAndDeleteIdCollection.finish(),\n        Search: Search.finish(),\n    }),\n    applyMiddleware(...middleware)\n)\n"
	//reg := regexp.MustCompile(`(\()(\s.+)combineReducers.+(\s.+)*}`)
	//fmt.Println(
	//	reg.FindString(content),
	//)

}
