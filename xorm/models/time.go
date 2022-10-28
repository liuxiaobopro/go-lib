package models

import (
	"database/sql/driver"
	"time"
)

// 源地址: https://www.jb51.net/article/202547.htm

const localDateTimeFormat string = "2006-01-02 15:04:05"

type LocalTime time.Time

func (l LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localDateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, localDateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (l *LocalTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+localDateTimeFormat+`"`, string(b), time.Local)
	*l = LocalTime(now)
	return err
}

func (l LocalTime) String() string {
	return time.Time(l).Format(localDateTimeFormat)
}

func (l LocalTime) Now() LocalTime {
	return LocalTime(time.Now())
}

func (l LocalTime) ParseTime(t time.Time) LocalTime {
	return LocalTime(t)
}

func (j LocalTime) format() string {
	return time.Time(j).Format(localDateTimeFormat)
}

func (j LocalTime) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}

// func (l *LocalTime) FromDB(b []byte) error {
// 	if nil == b || len(b) == 0 {
// 		l = nil
// 		return nil
// 	}
// 	var now time.Time
// 	var err error
// 	now, err = time.ParseInLocation(localDateTimeFormat, string(b), time.Local)
// 	if nil == err {
// 		*l = LocalTime(now)
// 		return nil
// 	}
// 	now, err = time.ParseInLocation("2006-01-02T15:04:05Z", string(b), time.Local)
// 	if nil == err {
// 		*l = LocalTime(now)
// 		return nil
// 	}
// 	// panic("自己定义个layout日期格式处理一下数据库里面的日期型数据解析!")
// 	return err
// }

// //func (t *LocalTime) Scan(v interface{}) error {
// // // Should be more strictly to check this type.
// // vt, err := time.Parse("2006-01-02 15:04:05", string(v.([]byte)))
// // if err != nil {
// // return err
// // }
// // *t = LocalTime(vt)
// // return nil
// //}

func (l *LocalTime) ToDB() ([]byte, error) {
	if nil == l {
		return nil, nil
	}
	return []byte(time.Time(*l).Format(localDateTimeFormat)), nil
}

func (l *LocalTime) Value() (driver.Value, error) {
	if nil == l {
		return nil, nil
	}
	return time.Time(*l).Format(localDateTimeFormat), nil
}
