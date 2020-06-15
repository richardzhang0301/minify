package min

import (
	"regexp"

	"github.com/richardzhang0301/minify"
	"github.com/richardzhang0301/minify/css"
	"github.com/richardzhang0301/minify/html"
	"github.com/richardzhang0301/minify/js"
	"github.com/richardzhang0301/minify/json"
	"github.com/richardzhang0301/minify/svg"
	"github.com/richardzhang0301/minify/xml"
)

// Default minifiers for CSS, HTML, XML, JS, JSON, and XML
var Default *minify.M

func init() {
	Default = minify.New()
	Default.AddFunc("text/css", css.Minify)
	Default.AddFunc("text/html", html.Minify)
	Default.AddFunc("image/svg+xml", svg.Minify)
	Default.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	Default.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	Default.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
}

// CSS string minifier using all default minifiers
func CSS(s string) (string, error) {
	return Default.String("text/css", s)
}

// HTML string minifier using all default minifiers
func HTML(s string) (string, error) {
	return Default.String("text/html", s)
}

// SVG string minifier using all default minifiers
func SVG(s string) (string, error) {
	return Default.String("image/svg+xml", s)
}

// JS string minifier using all default minifiers
func JS(s string) (string, error) {
	return Default.String("application/javascript", s)
}

// JSON string minifier using all default minifiers
func JSON(s string) (string, error) {
	return Default.String("application/json", s)
}

// XML string minifier using all default minifiers
func XML(s string) (string, error) {
	return Default.String("application/xml", s)
}