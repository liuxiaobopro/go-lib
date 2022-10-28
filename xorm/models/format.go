package models

import (
	"time"

	timel "gitee.com/liuxiaobopro/golib/utils/time"
)

// 将time.Time转换成自定义的time.Time
func FormatTime(t time.Time) time.Time {
	s := timel.TimeToString(t)
	// 将大写的T和+08:00替换成空格
	s = s[:10] + " " + s[11:19]
	// 将字符串转换成time.Time
	return timel.StringToTime(s)
}
