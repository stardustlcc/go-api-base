package timeutil

import (
	"math"
	"net/http"
	"time"
)

var (
	cst *time.Location
)

const CSTLayout = "2006-01-02 15:04:05"

func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}

	//默认设置为中国时区
	time.Local = cst
}

// 将 2020-11-08T08:18:46+08:00 => 2020-11-08 08:18:46
func RFC3339ToCSTLayout(value string) (string, error) {
	//将RFC3339的标准解析为 ts 类型
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	//使用 ts.In(cst) 将时间类型 ts 转换为指定时区
	//然后使用 Format 函数将时间格式化为 CSTLayout 所指定的字符串格式
	return ts.In(cst).Format(CSTLayout), nil
}

// 返回当前的时间日期格式
func CSTLayoutString() string {
	ts := time.Now()
	return ts.In(cst).Format(CSTLayout)
}

// 格式化当前日期格式的时间字符串
func ParseCSTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(CSTLayout, date, cst)
}

// 格式化为时间戳
func CSTLayoutStringToUnix(cstLayoutString string) (int64, error) {
	stamp, err := time.ParseInLocation(CSTLayout, cstLayoutString, cst)
	if err != nil {
		return 0, err
	}
	return stamp.Unix(), nil
}

// 返回 "Mon, 02 Jan 2006 15:04:05 GMT" 格式的时间
func GMTLayoutString() string {
	return time.Now().In(cst).Format(http.TimeFormat)
}

// 格式化GMT时间为当前
func ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, cst)
}

func SubInLocation(ts time.Time) float64 {
	return math.Abs(float64(time.Now().In(cst).Sub(ts).Seconds()))
}
