package utils

import "regexp"

const (
	// config 配置相关
	ConfigurationFileName = "package.json" // 配置文件名
	Profile_Field         = "conciseRedux" // 配置文件字段

	// error 报错相关
	configurationFileDoesNotExist                     = "the configuration file does not exist,please check package.json"                          // 配置文件不存在
	configurationFieldDoesNotExist                    = "the configuration field does not exist, please check if conciseRedux exists"              // 配置字段不存在
	FailedToParseConfigurationFile                    = "Failed to parse the configuration file because"                                           // 解析配置文件失败,因为....
	CanNotBeEmptyBaseURL                              = "baseURl can not be empty"                                                                 // baseURL 不能为空
	CanNotBeEmptyStorePath                            = "storePath can not be empty"                                                               // 不能为空 storeDirPath
	CanNotBeEmptyReducerDirPath                       = "reducerDirPath can not be empty"                                                          // 不能为空 reducerDirPath
	CanNotBeEmptyReducerList                          = "reducerList can not be empty"                                                             // 不能为空 reducerList
	StoreFileDoesNotExist                             = "store file does not exist"                                                                // store  文件不存在
	YouHaveIntroducedConsiderReduxButItIsNowCommented = "you have introduced ConsiderRedux but it is now commented"                                // 引入了 ConsiderRedux 但被注释
	PleaseCheckIfConsiderReduxIsIntroduced            = "please check if ConsiderRedux or combineReducers are introduced"                          // 请检查 是否引入 ConsiderRedux 或 combineReducers
	TheCombineReducersFunctionIsNotImplemented        = "the combineReducers function is not implemented"                                          // 没有实现  combineReducers  函数
	TheHooksFolderDoesNotExist                        = "the hooks folder does not exist, if you need to automatically create it, you can type -f" // hooks文件夹不存在，如需要自动创建可键入-f

	// generate

	GenerateAction                 = "Action"                            // 生成action name 常量
	GenerateActionTypes            = "ActionTypes"                       // 生成actionTypes name
	GenerateStateTypes             = "StateTypes"                        // 生成StateTypes
	GenerateActionPayloadTypes     = "ActionPayloadTypes"                // 生成actionPayloadTypes
	ReducerFileSuffix              = ".ts"                               // reducer文件后缀
	HooksFolderCreatedSuccessfully = "hooks folder created successfully" // hooks 文件夹创建成功
)

var (
	// regular Expression Judgment

	// TODO: 无法处理换行
	//WhetherToIntroduceConsideRedux   = regexp.MustCompile("(\\s+|.+)import(\\s+|)\\s.+combineReducer.+(from)\\s.+@zealforchange\\/conciseredux") // 是否引入 conciseredux
	WhetherToIntroduceConsideRedux   = regexp.MustCompile("combineReducer|@zealforchange\\/conciseredux") // 是否引入 conciseredux
	MatchFunctionBodyCombineReducers = regexp.MustCompile(`(\()(\s.+|)combineReducers.+(\s.+)*(\s|)}`)    // 匹配  combineReducers 函数体
	IsItComment                      = regexp.MustCompile("^(\\s.|)//")                                   // 是否是注释
)
