package erp

import (
	"github.com/seerwo/akspay/credential"
	"github.com/seerwo/akspay/erp/config"
	"github.com/seerwo/akspay/erp/context"
	"github.com/seerwo/akspay/erp/oauth"
	"github.com/seerwo/akspay/erp/order"
)


//Erp erp akspay api.
type Erp struct {
	ctx *context.Context
}

//NewPay instance pay api.
func NewErp(cfg *config.Config) *Erp {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyPayPrefix, cfg.Cache)
	ctx := &context.Context{
		Config: cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Erp{ctx: ctx}
}

//SetAccessTokenHandle custom access_token method.
func (erp *Erp) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle){
	erp.ctx.AccessTokenHandle = accessTokenHandle
}

//GetContext get Context
func (erp *Erp) GetContext() *context.Context {
	return erp.ctx
}

//GetAccessToken get access_token
func (erp *Erp) GetAccessToken()(string, error){
	return erp.ctx.GetAccessToken()
}

func (erp *Erp) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(erp.ctx)
}

//GetOrder get trade interface
func (erp *Erp) GetOrder() *order.Order {
	return order.NewOrder(erp.ctx)
}
