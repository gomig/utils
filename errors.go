package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TaggedError generate a tagged error
func TaggedError(tags []string, format string, args ...any) error {
	_tags := ""
	for _, t := range tags {
		_tags = fmt.Sprintf("%s[%s] ", _tags, t)
	}
	return fmt.Errorf(_tags+format, args...)
}

// IsErrorOf check if error has tag
func IsErrorOf(tag string, err error) bool {
	return strings.Contains(err.Error(), "["+tag+"]")
}

// HasError return true if error not nil, otherwise return false
func HasError(err error) bool {
	return err != nil
}

// PanicOnError generate panic if error is not null
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// ErrorOf get error of function result
func ErrorOf(res any, err error) error {
	return err
}

// VarOrPanic get function result (T, error)
//
// if result has error generate panic
// return T otherwise
func VarOrPanic[T any](res T, err error) T {
	PanicOnError(err)
	return res
}

// PrettyLog log data to output using json format
func PrettyLog(data any) {
	_bytes, _ := json.Marshal(data)
	fmt.Println(string(_bytes))
}

// PrettyLogIndent log data to output using json format with indent
func PrettyLogIndent(data any) {
	_bytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(_bytes))
}
