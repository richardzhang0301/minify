// +build gofuzz
package fuzz

import "github.com/richardzhang0301/minify"

func Fuzz(data []byte) int {
	m := minify.New()
	data = minify.DataURI(m, data)
	return 1
}
