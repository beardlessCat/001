package main

import (
	"fmt"
	"time"
)

/*
*
Unix Time是自1970年1月1日 00:00:00 UTC 至当前时间经过的总秒数。下面的代码片段演示了如何基于时间对象获取到Unix 时间。
*/
func main() {
	//时间戳
	now := time.Now()
	unix := now.Unix()       //秒级时间戳
	milli := now.UnixMilli() //毫秒级时间戳
	micro := now.UnixMicro() //微秒级时间戳
	nano := now.UnixNano()   //纳秒级时间戳
	println(unix, milli, micro, nano)
	//时间戳转时间
	// 获取北京时间所在的东八区时区对象
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	date := time.Date(2022, 12, 12, 12, 12, 12, 0, beijing)
	var (
		sec  = date.Unix()
		msec = date.UnixMilli()
		usec = date.UnixMicro()
	)
	// 将秒级时间戳转为时间对象（第二个参数为不足1秒的纳秒数）
	t := time.Unix(sec, 22)
	unixMilli := time.UnixMilli(msec)
	unixMicro := time.UnixMicro(usec)
	fmt.Println(t)
	fmt.Println(unixMilli)
	fmt.Println(unixMicro)
	fmt.Println("=======================================")
	//时间间隔
	//time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。time.Duration表示一段时间间隔，可表示的最长时间段大约290年。
	nowTime := time.Now()
	fmt.Println(nowTime)
	later := nowTime.Add(time.Hour)
	fmt.Println(later)
	eairly := nowTime.Sub(later)
	fmt.Println(eairly)
	equal := nowTime.Equal(later)
	before := nowTime.Before(later)
	after := nowTime.After(later)
	fmt.Println(equal)
	fmt.Println(before)
	fmt.Println(after)
	fmt.Println("==================================")
	//定时器
	//tick := time.Tick(time.Second)
	//for i := range tick {
	//	fmt.Println(i)
	//}
	//时间格式化
	format := nowTime.Format("2006-01-02 15:04:05.000 PM Mon Jan")
	fmt.Println(format)
}
