package pay

import (
	"github.com/seerwo/akspay/credential"
	"github.com/seerwo/akspay/pay/common"
	"github.com/seerwo/akspay/pay/config"
	"github.com/seerwo/akspay/pay/context"
	"github.com/seerwo/akspay/pay/oauth"
	"github.com/seerwo/akspay/pay/order"
)

//Pay pay akspay api.
type Pay struct {
	ctx *context.Context
}

//NewPay instance pay api.
func NewPay(cfg *config.Config) *Pay {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyPayPrefix, cfg.Cache)
	ctx := &context.Context{
		Config: cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Pay{ctx: ctx}
}

//SetAccessTokenHandle custom access_token method.
func (pay *Pay) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle){
	pay.ctx.AccessTokenHandle = accessTokenHandle
}

//GetContext get Context
func (pay *Pay) GetContext() *context.Context {
	return pay.ctx
}

//GetAccessToken get access_token
func (pay *Pay) GetAccessToken()(string, error){
	return pay.ctx.GetAccessToken()
}

func (pay *Pay) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(pay.ctx)
}

//GetOrder get trade interface
func (pay *Pay) GetOrder() *order.Order {
	return order.NewOrder(pay.ctx)
}

//GetCommon get common interface
func (pay *Pay) GetCommon() *common.Common {
	return common.NewCommon(pay.ctx)
}
