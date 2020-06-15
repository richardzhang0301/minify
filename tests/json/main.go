// +build gofuzz
package fuzz

import (
	"bytes"
	"io/ioutil"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/json"
)

func Fuzz(data []byte) int {
	r := bytes.NewBuffer(data)
	_ = json.Minify(minify.New(), ioutil.Discard, r, nil)
	return 1
}
