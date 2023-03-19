package types

type ConciseReduxConfig struct {
	ConciseRedux ConfigurationCollection `json:"conciseRedux"`
}

type ConfigurationCollection struct {
	BasicConfiguration
	AdvancedConfiguration
}
type ReducerListTypes = map[string][]string

type BasicConfiguration struct {
	BaseURL        string           `json:"baseURL"`
	StorePath      string           `json:"storePath"`      // 必填
	ReducerDirPath string           `json:"reducerDirPath"` // 必填
	ReducerList    ReducerListTypes `json:"reducerList"`    // 必填
	PathAlias      string           `json:"pathAlias"`      // 路径别名非必填
}

type AdvancedConfiguration struct {
	HooksDir string `json:"hooksDir"` // 生成自定义hooks路径
}

type InitializationProcess interface {
	CheckPath()                        // 检查路径
	CheckTheConfiguration(path string) // 检查配置内容
	ParseTheConfiguration()
}

type UpdateStore struct {
	UpdatedImport          []string // 新增的 import 路径
	AddReducerCall         []string // 新增 reducer 调用
	UpdatedCombineReducers string   // 更新后的 CombineReducers
}

type CmdConfig struct {
	CreateHooksFile bool // 创建和reducer 相关的hooks
	CheckHooks      bool // 检查hooks
	CheckReducer    bool // 检查reducer
}

type NewImportWithReducerOps struct {
	IntroduceStateStatement bool // 引入 对应的action的state声明？
}
