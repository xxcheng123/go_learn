module geecache

go 1.21.3

require github.com/pkg/errors v0.9.1

require lru v0.0.0
require (
	gee v0.0.0
)

replace lru => ./lru
replace gee => ./gee