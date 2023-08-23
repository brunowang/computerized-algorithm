package string

var digitMap = map[rune]int64{
	'一': 1, '二': 2, '三': 3, '四': 4, '五': 5,
	'六': 6, '七': 7, '八': 8, '九': 9, '零': 0,
}

var timesMap = map[rune]int64{
	'十': 1, '百': 2, '千': 3, '万': 4, '亿': 8,
}

func toDigit(str string) (res int64) {
	runes := []rune(str)
	var lastTimes int64
	var currTimes int64

	// 从后向前遍历字符串
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		// fmt.Println("origin", string(r))
		if v, ok := timesMap[r]; ok {
			if v > currTimes {
				// 如果进位标识大于累计进位，直接覆盖
				currTimes = v
			} else if v > lastTimes {
				// 如果进位标识大于上一个进位，则以当前进位标识为准
				currTimes += v - lastTimes
			} else {
				// 如果进位标识小于上一个进位，则继续累计进位
				currTimes += v
			}
			// 记录进位标识
			lastTimes = v
			// fmt.Println("times", currTimes)
			continue
		}
		if v, ok := digitMap[r]; ok {
			// 把当前数字放到合适的进位位置上
			res += v * powInt64(10, currTimes)
			// fmt.Println("digit", res)
		}
	}
	return
}

// 整数指数运算, 简单处理，不考虑负数
func powInt64(x, y int64) int64 {
	if y < 0 {
		return 0
	}
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}
	ret := x
	for i := int64(1); i < y; i++ {
		ret *= x
	}
	return ret
}
