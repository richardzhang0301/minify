// +build gofuzz
package fuzz

import (
	"github.com/richardzhang0301/minify/svg"
	"github.com/tdewolff/parse"
)

func Fuzz(data []byte) int {
	pathDataBuffer := svg.NewPathData(&svg.Minifier{Decimals: -1})
	data = parse.Copy(data) // ignore const-input error for OSS-Fuzz
	data = pathDataBuffer.ShortenPathData(data)
	return 1
}
