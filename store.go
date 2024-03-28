package errcode

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

var (
	langAlreadyExist = errors.New("lang already exist")
	store            = &localeStore{store: make(map[string]*locale)}
)

type (
	httpError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	errorCodes []*httpError

	locale struct {
		lan string
		Err errorCodes `json:"errCodes"`
	}

	localeStore struct {
		store map[string]*locale
	}
)

func from(lan, filename string) (*locale, error) {
	var (
		bytes []byte
		err   error
	)
	if bytes, err = os.ReadFile(filepath.Clean(filename)); err != nil {
		return nil, err
	}
	l := &locale{lan: lan}
	err = json.Unmarshal(bytes, &l)
	return l, err
}

func (ls *localeStore) addLocale(l *locale) bool {
	if _, ok := ls.store[l.lan]; ok {
		return false
	}

	ls.store[l.lan] = l
	return true
}

func (ls *localeStore) getFormat(lan string, format int) (string, bool) {
	if value, ok := ls.store[lan]; ok {
		for _, errCode := range value.Err {
			if errCode.Code == format {
				return errCode.Message, true
			}
		}
	}
	return strconv.Itoa(format), false
}
