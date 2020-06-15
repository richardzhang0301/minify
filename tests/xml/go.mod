module github.com/tdewolff/fuzz/minify/xml

go 1.13

replace github.com/richardzhang0301/minify => ../../../minify

require (
	github.com/dvyukov/go-fuzz v0.0.0-20191022152526-8cb203812681 // indirect
	github.com/richardzhang0301/minify v2.5.2
)
