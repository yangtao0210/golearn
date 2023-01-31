package singleton

import (
	"sync"
)

type singleton struct{}

var instance *singleton

// sync.Once  ->Do()保证某个操作仅且只执行一次
var once sync.Once

// GetInstance golang 单例模式
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
