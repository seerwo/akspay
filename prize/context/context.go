package context

import (
	"github.com/seerwo/akspay/credential"
	"github.com/seerwo/akspay/prize/config"
)

//Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
