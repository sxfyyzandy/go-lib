package starcache

import "time"

type ICache interface {
	Set(key string, value interface{})
	SetWithExp(key string, value interface{}, exp time.Duration)
	SetWithoutExp(key string, value interface{})

	Get(key string) (interface{}, bool)
	GetString(key string) (string, bool)

	Delete(key string)

	Purge()
}
