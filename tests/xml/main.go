// +build gofuzz
package fuzz

import (
	"bytes"
	"io/ioutil"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/xml"
)

func Fuzz(data []byte) int {
	r := bytes.NewBuffer(data)
	_ = xml.Minify(minify.New(), ioutil.Discard, r, nil)
	return 1
}
