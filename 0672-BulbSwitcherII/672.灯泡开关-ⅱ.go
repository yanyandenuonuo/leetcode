/*
 * @lc app=leetcode.cn id=672 lang=golang
 *
 * [672] 灯泡开关 Ⅱ
 */

// @lc code=start
func flipLights(n int, presses int) int {
	// solution1: 模拟
	// solution2: 找规律，按偶数次等于0次，所以只需要观察按奇数次的情况
	//			  开关1影响x
	//			  开关2影响2x
	//			  开关3影响2x+1
	//			  开关4影响3x+1
	//			  由上述观察可以知道x与x+6i状态一致
	//			  枚举6个灯泡状态
	//			  6i+1: 受开关1,3,4影响
	//			  6i+2: 受开关1,2影响
	//			  6i+3: 受开关1,3影响
	//			  6i+4: 受开关1,2,4影响
	//			  6i+5: 受开关1,3影响，与6i+3一致
	//			  6i+6: 受开关1,2影响，与6i+2一致
	//			  由此枚举1-4号灯泡状态即可，所以最多只有8种状态
	//			  此时可以写代码枚举1-4号状态
	statusMap := make(map[int]struct{}, 8)
	for bitIdx := 0; bitIdx < (1 << 4); bitIdx += 1 {
		oneCount, switchStatus := 0, [4]int{}
		for idx := 0; idx < 4; idx += 1 {
			switchStatus[idx] = (bitIdx >> idx) & 0x01
			oneCount += switchStatus[idx]
		}

		if oneCount&0x01 != presses&0x01 || oneCount > presses {
			// 确保奇偶性一致且按压次数有效
			continue
		}

		// 0^0 => 0, 0^1 => 1, 1^0 => 1, 1^1 => 0
		status := switchStatus[0] ^ switchStatus[2] ^ switchStatus[3]

		if n >= 2 {
			status |= ((switchStatus[0] ^ switchStatus[1]) << 1)
		}

		if n >= 3 {
			status |= ((switchStatus[0] ^ switchStatus[2]) << 2)
		}

		if n >= 4 {
			status |= ((switchStatus[0] ^ switchStatus[1] ^ switchStatus[3]) << 3)
		}

		statusMap[status] = struct{}{}
	}

	return len(statusMap)

	// solution3: 继续找规律
	//			  假设1-4开关分别按了a,b,c,d次，则灯泡状态为
	//			  6i+1 = (a+c+d)%2
	//			  6i+2 = (a+b)%2
	//			  6i+3 = (a+c)%2
	//			  6i+4 = (a+b+d)%2
	//			  6i+1和6i+3以及6i+2和6i+4均取决于d，所以1-3号灯的状态确定后可推断4的状态
	//			  由此观察1-3号灯状态即可
	//			  111
	//			  按0次 => [111]，有1种可能
	//			  按1次 => [000, 101, 010, 011]，有4种可能
	//			  按2次，有7种可能，缺少011
	//					000 => [111, 010, 101, 100]
	//					101 => [010, 111, 000, 001]
	//					010 => [101, 000, 111, 110]
	//					011 => [100, 001, 110, 111]
	//			  按3次，有8种可能，3个灯泡共有8种可能
	//			  		对按2次结果再按1次，因为存在开关1，一定存在按2次时同样的结果
	//			  		按任意2次再按1次，一定存在按1次时的结果，补全[011]
	//			  		将上述两种可能合并可获得8种可能结果
	/**
	if presses == 0 {
		return 1
	}

	switch n {
	case 1:
		// [0]00, [1]01, [0]10, [0]11
		return 2
	case 2:
		switch presses {
		case 1:
			// [00]0, [10]1, [01]0, [01]1
			return 3
		default:
			return 4
		}
	default:
		switch presses {
		case 1:
			return 4
		case 2:
			return 7
		default:
			return 8
		}
	}
	*/
}

// @lc code=end

