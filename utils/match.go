package utils

import (
	"fmt"
	"path/filepath"
	"reduxCli/types"
	"regexp"
	"unicode"
)

// MustImportConciseRedux 是否引入 ConsideRedux
func MustImportConciseRedux(content string) bool {
	result := WhetherToIntroduceConsideRedux.FindString(content)
	if isComment(result) {
		Red(YouHaveIntroducedConsiderReduxButItIsNowCommented)
		return false
	}
	return len(result) != 0
}

func isComment(content string) bool {
	result := IsItComment.FindString(content)
	return len(result) != 0

}

// MatchCombineReducers 匹配是否不为空
func MatchCombineReducers(content string) bool {
	return len(
		MatchFunctionBodyCombineReducers.FindString(content),
	) != 0
}

// CombineReducersImplementTheCorrespondingReducer CombineReducers是否实现实现对应reducers
func CombineReducersImplementTheCorrespondingReducer(content string, action string) bool {
	// (\s|)Common(\s|):(\s|)Common.finish\(\)
	reg := regexp.MustCompile(fmt.Sprint("(\\s|)", action, "(\\s|):(\\s|)", action, ".finish()"))
	return len(reg.FindString(content)) != 0
}

func NewImportWithReducer(config types.ConciseReduxConfig, actionName string) string {
	// 路径别名优先级最高
	var path string
	if len(config.ConciseRedux.PathAlias) != 0 {
		path = filepath.Join(config.ConciseRedux.PathAlias, config.ConciseRedux.ReducerDirPath, fmt.Sprint(actionName))
	} else {
		path = filepath.Join(config.ConciseRedux.BaseURL, config.ConciseRedux.ReducerDirPath, fmt.Sprint(actionName))
	}
	return fmt.Sprint("import { ", actionName, " } from '", path, "';", "\n")
}

func ReducerCall(actionName string) string {
	return fmt.Sprint(actionName, ":", actionName, ".finish(),\n")
}

func StringSliceToString(strArr []string) string {
	var result string
	for _, ele := range strArr {
		result += ele
	}
	return result
}

func UpdateCombineReducers(combineReducersContext string, updateReducerCall string) string {
	reg := regexp.MustCompile("}")
	return reg.ReplaceAllString(MatchFunctionBodyCombineReducers.FindString(combineReducersContext), fmt.Sprint(updateReducerCall, "}"))
}

// IsStartUpper 首字符是否大写
func IsStartUpper(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// CapitalizeTheFirstLetter 首字母大写
func CapitalizeTheFirstLetter(s string) string {
	if IsStartUpper(s) {
		return s
	}
	return Capitalize(s)
}
