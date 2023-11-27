package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/devfeel/dotweb/framework/crypto/des"
	"hash"
)

//ark sign method
const(
	SignTypeMD5 = `MD5`
	SignTypeHMACSHA256 = `HMAC-SHA256`
	SignTypeSHA256 = `SHA256`
)



//DES加密 ECB模式 无偏移量
func CalcuateAES(content, key1 string)(string, error){
	key := []byte(key1)
	plaintext := []byte(content)
	res, _ := des.ECBEncrypt(plaintext, key[:8])
	//fmt.Println()
	return hex.EncodeToString(res), nil
}

//CalculateSign calculate sign
func CalculateSign(content, signType, key string) (string, error){
	var h hash.Hash
	if signType == SignTypeHMACSHA256 {
		h = hmac.New(sha256.New, []byte(key))
	} else if signType == SignTypeSHA256 {
		h = sha256.New()
	} else {
		h = md5.New()
	}

	if _, err := h.Write([]byte(content)); err != nil {
		return ``, err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

//ParamSign calculate param sign
func ParamSign(p map[string]string, urlKey, key string) (string, error){
	str := OrderParam(p, urlKey, key)
	var signType string
	switch p["sign_type"]{
	case SignTypeMD5, SignTypeHMACSHA256, SignTypeSHA256 :
		signType = p["sign_type"]
	case ``:
		signType = SignTypeMD5
	default:
		return ``, errors.New(`invalid sign_type`)
	}
	return CalculateSign(str, signType, key)
}