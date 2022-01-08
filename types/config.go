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
}

type AdvancedConfiguration struct {
	HooksDir string `json:"hooksDir"` // 生成自定义hooks路径
}

type InitializationProcess interface {
	CheckPath()                        // 检查路径
	CheckTheConfiguration(path string) // 检查配置内容
	ParseTheConfiguration()
}
