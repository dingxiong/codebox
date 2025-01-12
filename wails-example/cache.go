package main

import (
	"github.com/hashicorp/golang-lru/v2"
)

var cache *lru.Cache[string, interface{}]


func GetOrCalculate[T any](key string, cal func(string) (*T, error) ) (*T, error) {
	value, ok := cache.Get(key)
	if ok {
		return value.(*T), nil
	}
	newValue, err := cal(key)
	if err != nil {
		return nil, err
	}
	cache.Add(key, newValue)
	return newValue, nil
}

func init() {
	var err error
	cache, err = lru.New[string, interface{}](1000)
	if err != nil {
		panic(err)
	}
}

