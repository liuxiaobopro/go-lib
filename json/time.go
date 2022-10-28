package json

import (
	"fmt"
	"time"

	"github.com/liuxiaobopro/go-lib/define"
)

type Time time.Time

var NilTime = Time{}
var TimeFormats = []string{define.TimeFormat, define.TimeFormatStr}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	fmt.Println(string(data))
	// 空值不进行解析
	if len(data) == 2 {
		*t = Time(time.Time{})
		return
	}

	var now time.Time
	for _, format := range TimeFormats {
		// 指定解析的格式
		if now, err = time.ParseInLocation(format, string(data), time.Local); err == nil {
			*t = Time(now)
			return
		}
		// 指定解析的格式
		if now, err = time.ParseInLocation(`"`+format+`"`, string(data), time.Local); err == nil {
			*t = Time(now)
			return
		}
	}

	return
}
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormats[0])+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormats[0])
	b = append(b, '"')
	return b, nil
}
func (t Time) String() string {
	return time.Time(t).Format(TimeFormats[0])
}
