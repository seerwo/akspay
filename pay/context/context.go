package context

import (
	"github.com/seerwo/akspay/credential"
	"github.com/seerwo/akspay/pay/config"
)

//Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
