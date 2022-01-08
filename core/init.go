package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reduxCli/template"
	"reduxCli/types"
	"reduxCli/utils"
)

type InitializationProcess struct {
	PathName               string
	configPath             string                   // 配置文件路径
	configurationStructure types.ConciseReduxConfig // 配置结构体
}

func (that *InitializationProcess) CheckPath() *InitializationProcess {
	URL, _, readURLErr := utils.CheckConfigurationFileExist(that.PathName)
	// 如果不存在package.json文件就报错
	_ = utils.CatchErr(readURLErr, func(msg string) string {
		return msg
	}, true)
	that.configPath = URL
	return that
}

func (that *InitializationProcess) CheckTheConfiguration() *InitializationProcess {

	file, fileCloseHandler := utils.OpenFile(that.configPath)
	defer func() {
		fileCloseHandler()
	}()
	byteFile, _ := ioutil.ReadAll(file)
	mappingConfiguration := types.ConciseReduxConfig{}
	parseErr := json.Unmarshal(byteFile, &mappingConfiguration)
	_ = utils.CatchErr(parseErr, func(msg string) string {
		return msg
	}, true)
	utils.CheckNull(
		mappingConfiguration.ConciseRedux.BaseURL,
		utils.CanNotBeEmptyBaseURL,
	)
	utils.CheckNull(
		mappingConfiguration.ConciseRedux.StorePath,
		utils.CanNotBeEmptyStorePath,
	)
	utils.CheckNull(
		mappingConfiguration.ConciseRedux.ReducerDirPath,
		utils.CanNotBeEmptyReducerDirPath,
	)
	utils.CheckNull(
		mappingConfiguration.ConciseRedux.ReducerList,
		utils.CanNotBeEmptyReducerList,
	)
	that.configurationStructure = mappingConfiguration
	return that
}
func (that *InitializationProcess) generateReducer() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, that.configurationStructure.ConciseRedux.BaseURL, that.configurationStructure.ConciseRedux.ReducerDirPath)
	if !utils.IsExist(path) {
		utils.MkDir(path)
	}
	for reducerName, actionsList := range that.configurationStructure.ConciseRedux.ReducerList {
		reducerPath := fmt.Sprint(
			filepath.Join(path, reducerName), utils.ReducerFileSuffix,
		)
		if !utils.IsExist(reducerPath) {
			go utils.CreateFile(reducerPath,
				template.CreateReducer(reducerName, actionsList),
			)
		}
	}
}
func (that *InitializationProcess) updateStoreFile() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, that.configurationStructure.ConciseRedux.BaseURL, that.configurationStructure.ConciseRedux.StorePath)
	if !utils.IsExist(path) {
		panic(errors.New(utils.StoreFileDoesNotExist))
	}
	fileFragmentation := utils.ReadFile(path)
	fileContent := string(fileFragmentation)
	// 未引入包 或 combineReducers 函数
	if !utils.MustImportConciseRedux(fileContent) {
		panic(errors.New(utils.PleaseCheckIfConsiderReduxIsIntroduced))
	}
	// 没有实现 combineReducers 函数
	if !utils.MatchCombineReducers(fileContent) {
		panic(errors.New(utils.TheCombineReducersFunctionIsNotImplemented))
	}
}

func (that *InitializationProcess) ParseTheConfiguration() {
	that.generateReducer()
	that.updateStoreFile()
}
