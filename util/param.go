package util

import (
	"bytes"
	"sort"
)

//OrderParam trade params
func OrderParam(p map[string]string, bizKey string, key string) (returnStr string){
	keys := make([]string, 0, len(p))
	key = "&" + key
	for k := range p {
		if k == "sign_type" {
			continue
		}
		keys = append(keys, k)
	}
	var buf bytes.Buffer
	for _, k := range keys {
		if p[k] == ""{
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(p[k])
	}
	if buf.Len() > 0 {

		returnStr += bizKey + "&"
	}
	buf.WriteString(key)
	returnStr +=  buf.String()
	return
}

func UrlParam(p map[string]string, bizKey string) (returnStr string) {
	keys := make([]string, 0, len(p))
	for k := range p {
		if k == "sign_type" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if p[k] == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(p[k])
	}
	if buf.Len() > 0 {
		returnStr += bizKey + "?"
	}
	returnStr += buf.String()
	return
}
