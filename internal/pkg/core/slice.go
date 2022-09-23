package core

import (
	"errors"
	"reflect"
)

func Segment(a []interface{}, chuckSize int) ([][]interface{}, error) {
	if chuckSize < 1 {
		return nil, errors.New("chuckSize must be greater that zero")
	}
	chunks := make([][]interface{}, 0, (len(a)+chuckSize-1)/chuckSize)

	for chuckSize < len(a) {
		a, chunks = a[chuckSize:], append(chunks, a[0:chuckSize:chuckSize])
	}
	chunks = append(chunks, a)
	return chunks, nil
}

func ToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
