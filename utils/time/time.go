package time

import (
	"time"

	"github.com/liuxiaobopro/go-lib/define"
	jsonl "github.com/liuxiaobopro/go-lib/json"
)

var TimeFormat = define.TimeFormat
var TimeFormatStr = define.TimeFormatStr

// GetNowTime 获取当前时间
func GetNowTime() string {
	return time.Now().Format(TimeFormat)
}

// GetNowTimeUnsigned 获取当前时间
func GetNowTimeUnsigned() string {
	return time.Now().Format(TimeFormatStr)
}

// GetNowDate 获取当前日期
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// GetNowTimeUnix 获取当前时间戳
func GetNowTimeUnix() int64 {
	return time.Now().Unix()
}

// GetNowTimeUnixMilli 获取当前时间戳(毫秒)
func GetNowTimeUnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetNowTimeUnixNano 获取当前时间戳(纳秒)
func GetNowTimeUnixNano() int64 {
	return time.Now().UnixNano()
}

// DateToUnix 日期转时间戳
func DateToUnix(date string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, date, loc)
	sr := theTime.Unix()
	return sr
}

// UnixToDate 时间戳转日期
func UnixToDate(timestamp int64) string {
	timeLayout := TimeFormat //转化所需模板
	dataTimeStr := time.Unix(timestamp, 0).Format(timeLayout)
	return dataTimeStr
}

// UnixToDateMilli 时间戳(毫秒)转日期
func UnixToDateMilli(timestamp int64) string {
	timeLayout := TimeFormat //转化所需模板
	dataTimeStr := time.Unix(timestamp/1000, 0).Format(timeLayout)
	return dataTimeStr
}

// UnixToDateNano 时间戳(纳秒)转日期
func UnixToDateNano(timestamp int64) string {
	timeLayout := TimeFormat //转化所需模板
	dataTimeStr := time.Unix(0, timestamp).Format(timeLayout)
	return dataTimeStr
}

// GetTimeUnix 获取指定时间戳
func GetTimeUnix(year, month, day, hour, min, sec int) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, time.Date(year, time.Month(month), day, hour, min, sec, 0, loc).Format(TimeFormat), loc)
	sr := theTime.Unix()
	return sr
}

// GetTimeUnixMilli 获取指定时间戳(毫秒)
func GetTimeUnixMilli(year, month, day, hour, min, sec int) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, time.Date(year, time.Month(month), day, hour, min, sec, 0, loc).Format(TimeFormat), loc)
	sr := theTime.UnixNano() / 1e6
	return sr
}

// GetTimeUnixNano 获取指定时间戳(纳秒)
func GetTimeUnixNano(year, month, day, hour, min, sec int) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, time.Date(year, time.Month(month), day, hour, min, sec, 0, loc).Format(TimeFormat), loc)
	sr := theTime.UnixNano()
	return sr
}

// GetTimeUnixByDate 获取指定时间戳
func GetTimeUnixByDate(date string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, date, loc)
	sr := theTime.Unix()
	return sr
}

// GetTimeUnixMilliByDate 获取指定时间戳(毫秒)
func GetTimeUnixMilliByDate(date string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, date, loc)
	sr := theTime.UnixNano() / 1e6
	return sr
}

// GetTimeUnixNanoByDate 获取指定时间戳(纳秒)
func GetTimeUnixNanoByDate(date string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, date, loc)
	sr := theTime.UnixNano()
	return sr
}

// StringToTime string转time.Time
func StringToTime(str string) time.Time {
	t, _ := time.Parse(TimeFormat, str)
	return t
}

// TimeToString time.Time转string
func TimeToString(t time.Time) string {
	return t.Format(TimeFormat)
}

func JsonlTimeToString(t jsonl.Time) string {
	return time.Time(t).Format(TimeFormat)
}

func StringToJsonlTime(str string) jsonl.Time {
	t, _ := time.Parse(TimeFormat, str)
	return jsonl.Time(t)
}
