package util

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

//RandomStr generate string
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Interface2String(inter interface{})(str string) {

	switch inter.(type) {
	case string:
		str = inter .(string)
		break
	case int:
		str = strconv.Itoa(inter.(int))
		break
	case int64:
		str = strconv.FormatInt(inter.(int64), 10)
		break
	case float64:
		str =  strconv.FormatFloat(inter.(float64), 'f',2, 64)
		break
	default:
		b, _ := json.Marshal(inter)
		str = string(b)
	}

	return
}