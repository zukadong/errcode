package errcode

import (
	"fmt"
	"testing"
)

const (
	zhCN = "zh-CN"
	enUS = "en-US"
)

func init() {
	for _, lan := range []string{zhCN, enUS} {
		if err := TryLoadErrCodeConfig(lan, "testdata/errcode_"+lan+".json"); err != nil {
			panic("load errcode config failed, error: " + err.Error())
		}
	}
}

func TestErrorCode(t *testing.T) {
	// direct invoke
	fmt.Println(GetErrMessage(enUS, 10002, "World!"))
	fmt.Println(GetErrMessage(zhCN, 10002, "世界!"))

	// wrap Locale
	fmt.Println((&Locale{Lan: enUS}).GetErrMessage(10002, "World!"))
	fmt.Println((&Locale{Lan: zhCN}).GetErrMessage(10002, "世界!"))
}
