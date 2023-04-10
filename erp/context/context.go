package context

import (
	"github.com/seerwo/akspay/credential"
	"github.com/seerwo/akspay/erp/config"
)

//Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
