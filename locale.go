package errcode

import (
	"fmt"
	"reflect"
	"strconv"
)

func TryLoadErrCodeConfig(lan, filename string) error {
	l, err := from(lan, filename)
	if err != nil {
		return err
	}
	if !store.addLocale(l) {
		return langAlreadyExist
	}
	return nil
}

type Locale struct {
	Lan string
}

func (l Locale) GetErrMessage(format int, args ...any) string {
	return GetErrMessage(l.Lan, format, args...)
}

func GetErrMessage(lan string, format int, args ...any) string {
	value, ok := store.getFormat(lan, format)
	if ok {
		if len(args) > 0 {
			//return fmt.Sprintf(value, parseParams(args...)...)
			return fmt.Sprintf(value, args...)
		}
		return value
	}
	return strconv.Itoa(format)
}

func parseParams(args ...any) []any {
	params := make([]any, 0, len(args))
	for _, arg := range args {
		if arg == nil {
			continue
		}

		val := reflect.ValueOf(arg)
		if val.Kind() == reflect.Slice {
			for i := 0; i < val.Len(); i++ {
				tmp := val.Index(i).Interface()
				if len(tmp.(string)) == 0 {
					continue
				}
				params = append(params, tmp)
			}
		} else {
			params = append(params, arg)
		}
	}
	return params
}
