/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	isMinus := false
	if numerator < 0 {
		isMinus = !isMinus
		numerator = -numerator
	}
	if denominator < 0 {
		isMinus = !isMinus
		denominator = -denominator
	}
	intVal := numerator / denominator
	remainderVal := numerator % denominator
	remainderValMap := make(map[int]int)

	fractionalList := make([]byte, 0)
	cyclicIdx := 0

	for idx := 0; remainderVal > 0; idx += 1 {
		if remainderVal < denominator {
			remainderValMap[remainderVal] = idx
			remainderVal *= 10
			fractionalList = append(fractionalList, '0')
			continue
		}
		fractionalList = append(fractionalList, byte('0'+remainderVal/denominator))
		remainderVal = remainderVal % denominator

		if valIdx, isExist := remainderValMap[remainderVal]; isExist {
			// 找到循环位置
			cyclicIdx = valIdx
			break
		}
		remainderValMap[remainderVal] = idx

		remainderVal *= 10
	}

	res := strconv.Itoa(intVal)
	if remainderVal == 0 {
		// 非循环小数
		if len(fractionalList) > 0 {
			res += "." + string(fractionalList[1:])
		}
	} else {
		// 循环小数
		res += "."
		if cyclicIdx > 0 {
			res += string(fractionalList[1 : cyclicIdx+1])
		}
		res += "(" + string(fractionalList[cyclicIdx+1:]) + ")"
	}
	if isMinus && (intVal > 0 || len(fractionalList) > 0) {
		res = "-" + res
	}
	return res
}

// @lc code=end

