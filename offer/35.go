package offer

import (
	"algorithm/utils"
	"strconv"
	"strings"
)

//面试题35：最小时间差
//题目：给定一组范围在00：00至23：59的时间，求任意两个时间之间的最小时间差。
//例如，输入时间数组["23：50"，"23：59"，"00：00"]，"23：59"和"00：00"之间只有1分钟的间隔，是最小的时间差。

func FindMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}

	minuteFlag := make([]bool, 1440)
	for _, time := range timePoints {
		s := strings.Split(time, ":")
		hour, _ := strconv.ParseInt(s[0], 10, 64)
		min, _ := strconv.ParseInt(s[1], 10, 64)
		min += 60 * hour
		if minuteFlag[min] {
			return 0
		}
		minuteFlag[min] = true
	}
	return findMinDifferenceHelper(minuteFlag)
}

func findMinDifferenceHelper(minuteFlag []bool) int {
	var (
		minDiff = len(minuteFlag) - 1
		prev    = -1
		first   = len(minuteFlag) - 1
		last    = -1
	)
	for i := 0; i < len(minuteFlag); i++ {
		if minuteFlag[i] {
			if prev >= 0 {
				minDiff = utils.Min(minDiff, i-prev)
			}
			prev = i
			first = utils.Min(first, i)
			last = utils.Max(last, i)
		}
	}

	minDiff = utils.Min(first+len(minuteFlag)-last, minDiff)
	return minDiff
}
