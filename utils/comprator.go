package utils

import (
	"strings"
	"bytes"
	"fmt"
)

type Comprator func(a, b interface{}) (int, error)

func StringComprator(a, b interface{}) (int, error) {
	sa,ok := a.(string)
	if !ok {
		return 0, fmt.Errorf("a is not string.")
	}
	sb, ok := b.(string)
	if !ok {
		return 0, fmt.Errorf("b is not string.")
	}
	return strings.Compare(sa, sb), nil
}

func byteArrayComprator(a, b interface{}) int {
	ba := a.([]byte)
	bb := b.([]byte)
	return bytes.Compare(ba, bb)
}