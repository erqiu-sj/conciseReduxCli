package utils

import (
	"errors"
)

// CatchErr 捕获错误
func CatchErr(err error, failCb func(msg string) string, triggerPanic bool) error {

	if err != nil {
		errMsg := failCb(err.Error())
		if triggerPanic {
			panic(errors.New(errMsg))
		}
	}
	return err
}

func CheckNull(val interface{}, panicMsg string) {
	switch val.(type) {
	case string:
		if val == "" {
			panic(errors.New(panicMsg))
		}
		return
	case map[string][]string:
		assertionSuc, _ := val.(map[string][]string)
		if len(assertionSuc) == 0 {
			panic(errors.New(panicMsg))
		}
		return
	default:
		return
	}
}
