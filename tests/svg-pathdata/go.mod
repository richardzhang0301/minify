module github.com/tdewolff/fuzz/minify/svg-pathdata

go 1.13

replace github.com/richardzhang0301/minify => ../../../minify

require (
	github.com/dvyukov/go-fuzz v0.0.0-20191022152526-8cb203812681 // indirect
	github.com/richardzhang0301/minify v2.5.2
	github.com/tdewolff/parse v2.3.4+incompatible
)
