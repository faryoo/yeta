package context

import (
	"github.com/faryoo/yeta/credential"
	"github.com/faryoo/yeta/work/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
