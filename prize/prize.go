package prize

import (
	"github.com/seerwo/akspay/credential"
	"github.com/seerwo/akspay/prize/config"
	"github.com/seerwo/akspay/prize/context"
	"github.com/seerwo/akspay/prize/oauth"
	"github.com/seerwo/akspay/prize/user"
)

//Prize prize akspay api.
type Prize struct {
	ctx *context.Context
}

//NewPrize instance prize api.
func NewPrize(cfg *config.Config) *Prize {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyPayPrefix, cfg.Cache)
	ctx := &context.Context{
		Config: cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Prize{ctx: ctx}
}

//SetAccessTokenHandle custom access_token method.
func (prize *Prize) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle){
	prize.ctx.AccessTokenHandle = accessTokenHandle
}

//GetContext get Context
func (prize *Prize) GetContext() *context.Context {
	return prize.ctx
}

//GetAccessToken get access_token
func (prize *Prize) GetAccessToken()(string, error){
	return prize.ctx.GetAccessToken()
}

func (prize *Prize) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(prize.ctx)
}

//GetUser get trade interface
func (prize *Prize) GetUser() *user.User {
	return user.NewUser(prize.ctx)
}