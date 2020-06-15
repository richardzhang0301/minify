// +build gofuzz
package fuzz

import (
	"github.com/richardzhang0301/minify"
	"github.com/tdewolff/parse"
)

func Fuzz(data []byte) int {
	data = parse.Copy(data) // ignore const-input error for OSS-Fuzz
	data = minify.Mediatype(data)
	return 1
}
