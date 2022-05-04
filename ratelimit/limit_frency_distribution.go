package main

import (
	"strconv"
	"time"
)

func LimitFreqs(queueName string, count uint, timeWindow int64) bool {
	currTime := time.Now().Unix()
	length := uint(ListLen(queueName))
	if length < count {
		ListPush(queueName, currTime)
		return true
	}
	//队列满了,取出最早访问的时间
	earlyTime, _ := strconv.ParseInt(ListIndex(queueName, int64(length)-1), 10, 64)
	//说明最早期的时间还在时间窗口内,还没过期,所以不允许通过
	if currTime-earlyTime <= timeWindow {
		return false
	} else {
		//说明最早期的访问应该过期了,去掉最早期的
		ListPop(queueName)
		ListPush(queueName, currTime)
	}
	return true
}
