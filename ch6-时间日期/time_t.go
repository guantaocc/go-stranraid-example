package main

import (
	"fmt"
	"time"
)

// 日期时间处理函数

// 主要类型
// Location: 地理位置偏移, 默认UTC时区
// Time: 纳秒精度时间点
// Duration: 两个时间点经过的时间
// Timer 和 Ticker：定时器相关类型
// Weekday 和 Month:

// 相关方法
// time.Unix(sec, nsec int64) 通过 Unix 时间戳生成 time.Time 实例；
// time.Time.Unix() 得到 Unix 时间戳；
// time.Time.UnixNano() 得到 Unix 时间戳的纳秒表示；

// 格式化和解析
// time.Parse, time.ParseLocation, time.Time.Format()

func main() {
	fmt.Println(time.Time.Unix(time.Now()))
	fmt.Println(time.Time.UnixNano(time.Now()))
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 09:14:00", time.Local)
	fmt.Println(time.Now().Sub(t).Hours())
}
