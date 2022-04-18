package main

import (
	"fmt"
	"time"
)

func main() {
	diff := hourDiffer("2020-08-10 14:17:57.420", "2020-08-10 14:17:58.420")
	fmt.Print(diff)
	timeUnix := time.Now().Unix()         //单位秒
	timeUnixNano := time.Now().UnixNano() //单位纳秒
	fmt.Println(timeUnix)
	fmt.Println(timeUnixNano)

	datetime := "2020-08-11 10:20:18.239" //待转化为时间戳的字符串

	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.UnixNano() //转化为时间戳 类型是int64
	fmt.Println(timestamp)

	endTime := "2020-08-11 10:20:20.339"
	//日期转化为时间戳
	tmp1, _ := time.ParseInLocation(timeLayout, endTime, loc)
	tmp1s := tmp1.UnixNano() //转化为时间戳 类型是int64
	fmt.Println(tmp1s)

	fmt.Printf("diff:%d\n", (tmp1s-timestamp)/1000000)

	//时间戳转化为日期
	datetime = time.Unix(timestamp, 0).Format(timeLayout)
	fmt.Println(datetime)

}
func hourDiffer(start_time string, end_time string) int64 {
	var hour int64
	t1, _ := time.ParseInLocation("yyyy-MM-dd hh24:mi:ss.DDD", start_time, time.Local)
	t2, _ := time.ParseInLocation("yyyy-MM-dd hh24:mi:ss.DDD", end_time, time.Local)
	if t2.Before(t1) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff
		return hour
	} else {
		return hour
	}
}
