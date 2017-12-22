package sqlgen

import (
	"github.com/ilovetravel/upperiodb/util/cache"
)

type cc interface {
	cache.Cacheable
	compilable
}

type compilable interface {
	Compile(*Template) string
}
