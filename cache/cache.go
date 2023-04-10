package cache

import "time"

//Cache interface
type Cache interface {
	Get(Key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
}
