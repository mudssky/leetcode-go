package leetcode

func lemonadeChange(bills []int) bool {
	// five, ten, twenty := 0, 0, 0
	five, ten := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		}
		if bills[i] == 10 {
			if five <= 0 {
				return false
			}
			ten++
			five--
		}
		if bills[i] == 20 {
			// 1.优先消耗10美元，因为5美元找零更好用
			if five > 0 && ten > 0 {
				five--
				ten--
				// twenty++//因为20不会用来找零，所以这里也可以不记录，删掉20这个变量也可以
			} else if five >= 3 {
				five -= 3
				// twenty++
			} else {
				return false
			}

		}
	}
	return true
}
