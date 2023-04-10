package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/seerwo/akspay/prize/context"
	"github.com/seerwo/akspay/util"
)

const (
	USER_URL = "/sign/api/openapi_userVerify"
)

//User user struct
type User struct {
	*context.Context
}

type ReqUser struct{
	Tel string `json:"tel"`
}

type ResUser struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

//NewPackage instance
func NewUser(context *context.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

func (u *User) GetUserReq(req ReqUser)(res ResUser, err error) {

	//accessParam, _ := json.Marshal(req)
	var accessParam string
	if accessParam, err = u.GetPrizeParam(req.Tel); err != nil {
		return
	}
	//fmt.Println(accessUrl)
	//fmt.Println(string(accessParam))
	var response []byte
	url := fmt.Sprintf(util.BASE_PRIZE_WEB_URL + USER_URL)
	//fmt.Println(url)
	if response, err = util.HTTPPost(url, accessParam); err != nil {
		return
	}

	response = bytes.TrimPrefix(response, []byte("\xef\xbb\xbf"))
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	if res.Code != 1 {
		err = fmt.Errorf("GetUserReq Error , errcode=%s , errmsg=%s", res.Code, res.Msg)
		return
	}
	return
}