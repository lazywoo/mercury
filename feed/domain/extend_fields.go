package domain

import "errors"

type ExtendFields map[string]string

var errKeyNotFound = errors.New("key not found")

func (f ExtendFields) Get(key string) (string, error) {
	val, ok := f[key]
	if !ok {
		return "", errKeyNotFound
	}
	return val, nil
}
