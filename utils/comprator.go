package utils

import (
	"strings"
	"bytes"
)

type Comprator func(a, b interface{}) int

func StringComprator(a, b interface{}) int {
	sa := a.(string)
	sb := b.(string)
	return strings.Compare(sa, sb)
}

func byteArrayComprator(a, b interface{}) int {
	ba := a.([]byte)
	bb := b.([]byte)
	return bytes.Compare(ba, bb)
}