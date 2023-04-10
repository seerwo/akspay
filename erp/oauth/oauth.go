package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/akspay/erp/context"
	"github.com/seerwo/akspay/util"
	"net/http"
	"net/url"
)

const(
	redirectOauthURL = "%s%s%s%s"
	webAppRedirectOauthURL = "%s%s%s%s"
	accessTokenURL = "%s%s%s"
	refreshAccessTokenURL = "%s%s"
	userInfoURL = ""
	checkAccessTokenURL = "%s%s"
)

//Oauth save user oauth message
type Oauth struct {
	*context.Context
}

//NewOauth instance oauth message
func NewOauth(context *context.Context) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	return auth
}

//GetRedirectURL to jump the url.
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, state string) (string, error){
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(redirectOauthURL, oauth.AppID, urlStr, scope, state), nil
}

//GetWebAppRedirectURL get web application to jump the url
func (oauth *Oauth) GetWebAppRedirectURL(redirectURI, scope, state string) (string, error){
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(webAppRedirectOauthURL, oauth.AppID, urlStr, scope, state), nil
}

//Redirect to jump the oauth
func (oauth *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, redirectURI, scope, state string) error {
	location, err := oauth.GetRedirectURL(redirectURI, scope, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, http.StatusFound)
	return nil
}

//ResAccessToken to get user oauth access_token result
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID string `json:"openid"`
	Scope string `json:"scope"`
}

//GetUserAccessToken to get oauth code to exchange access_token
func (oauth *Oauth) GetUserAccessToken(code string)(result ResAccessToken, err error){
	urlStr := fmt.Sprintf(accessTokenURL, oauth.AppID, oauth.AppSecret, code)
	var response []byte
	if response, err = util.HTTPGet(urlStr); err != nil {
		return
	}
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.ErrCode != "0" {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//RefreshAccessToken refresh access_token
func (oauth *Oauth) RefreshAccessToken(refreshToken string) (result ResAccessToken, err error){
	urlStr := fmt.Sprintf(refreshAccessTokenURL, oauth.AppID, refreshToken)
	var response []byte
	if response, err = util.HTTPGet(urlStr); err != nil {
		return
	}
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.ErrCode  != "0" {
		err = fmt.Errorf("GetUserAcessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//CheckAccessToken to check access_token
func (oauth *Oauth) CheckAccessToken(accessToken, openID string)(b bool, err error){
	urlStr := fmt.Sprintf(checkAccessTokenURL, accessToken, openID)
	var response []byte
	if response, err = util.HTTPGet(urlStr); err != nil {
		return
	}
	var result util.CommonError
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.ErrCode != "0" {
		b = false
		return
	}
	b = true
	return
}