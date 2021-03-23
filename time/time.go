// Package time 公共组件
package time

import (
	"time"
)

// GetSec 返回现在的Unix秒时间戳
func GetSec() int64 {
	return time.Now().Unix()
}

// GetMilSec 返回当前的UNIX毫秒时间戳
func GetMilSec() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetMicroSec 返回当前UNIX微秒时间戳
func GetMicroSec() int64 {
	return (time.Now().UnixNano() / 1000)
}

// GetBigSec 返回50年之后的时间戳
func GetBigSec() int64 {
	return (GetSec() + 50*365*24*60*60)
}

// GetDay 返回 day
func GetDay() int {
	return time.Now().Day()
}

// GetDate 返回year, month, day
func GetDate() (int, int, int) {
	return time.Now().Year(), int(time.Now().Month()), time.Now().Day()
}
