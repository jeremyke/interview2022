package main

import (
	"sort"
)

type meetingTime struct {
	start int
	end   int
}

type meetingTimeArr []*meetingTime

func (array meetingTimeArr) Len() int {
	return len(array)
}

func (array meetingTimeArr) Less(i, j int) bool {
	return array[i].end < array[j].end //从小到大， 若为大于号，则从大到小
}

func (array meetingTimeArr) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

//返回可以安排的会议场数
func bestArrange(aMeetingTimeArr []*meetingTime, timePoint int) int {
	sort.Sort(meetingTimeArr(aMeetingTimeArr)) //按照结束时间从小到大排序
	result := 0
	for i := 0; i < len(aMeetingTimeArr); i++ {
		if timePoint <= aMeetingTimeArr[i].start {
			result++
			timePoint = aMeetingTimeArr[i].end
		}
	}
	return result
}
