package credential

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/akspay/cache"
	"github.com/seerwo/akspay/util"
	"sync"
	"time"
)

const(
	//AccessTokenURL get access_token interface
	acccessTokenURL = ""
	//CacheKeyOpenPrefix ark cache key prefix
	CacheKeyPayPrefix="go_akspay_pay_"
)

//DefaultAccessToken default get AccessToken
type DefaultAccessToken struct {
	appID string
	appSecret string
	cacheKeyPrefix string
	cache cache.Cache
	accessTokenLock *sync.Mutex
}

//NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(appID, appSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil{
		panic("cache is ineed")
	}
	return &DefaultAccessToken{
		appID: appID,
		appSecret: appSecret,
		cache: cache,
		cacheKeyPrefix: cacheKeyPrefix,
		accessTokenLock:new(sync.Mutex),
	}
}

//ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
}

//GetAccessToken get access_token from cache value.
func (ak *DefaultAccessToken) GetAccessToken()(accessToken string, err error){
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.appID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	//cache invalid, reforce from server.
	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(ak.appID, ak.appSecret)
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires) * time.Second)
	if err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

func (ak *DefaultAccessToken) GetPrizeParam(tel string)(accessParam string, err error){


	//accessParam= fmt.Sprintf(url, ak.appID, clientReportId)
	m := make(map[string]string)
	m["tel"], err = util.CalcuateAES(tel, ak.appID)
	//m["sign"] = util.RandomStr(16)
	//m["sign_type"] = "SHA256"
	//
	m["sign"], err = util.CalculateSign(tel + ak.appSecret, "","")
	//accessParam = util.UrlParam(m, accessParam)
	jsonStr, err := json.Marshal(m)
	accessParam = string(jsonStr)

	fmt.Println(accessParam)
	return
}

func (ak *DefaultAccessToken) GetAccessParam(req interface{}, url, clientReportId string)(accessParam string, err error){

	accessParam= fmt.Sprintf(url, ak.appID, clientReportId)
	m := make(map[string]string)
	m["time"] = util.GetCurrTSStr()
	m["nonce_str"] = util.RandomStr(16)
	m["sign_type"] = "SHA256"

	m["sign"], err = util.ParamSign(m, ak.appID, ak.appSecret)
	accessParam = util.UrlParam(m, accessParam)

	return
}


func GetTokenFromServer(appID, appSecret string) (resAccessToken ResAccessToken, err error){
	url := fmt.Sprintf("%s?", appID)
	var body []byte
	if body, err = util.HTTPGet(url); err != nil {
		return
	}
	if err = json.Unmarshal(body, &resAccessToken); err != nil {
		return
	}
	if resAccessToken.ErrMsg != ""{
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}
