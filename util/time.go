package util

import (
	"strconv"
	"time"
)

//GetCurrTS return current timestamps
func GetCurrTS() int64 {
	return time.Now().UnixNano() / 1e6
}

//GetCurrTSStr return current timestamps
func GetCurrTSStr() string {
	return strconv.FormatInt(GetCurrTS(),10)
}
