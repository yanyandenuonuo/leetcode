/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	for idx := 0; idx < len(intervals); idx += 1 {
		interval := intervals[idx]
		if newInterval[1] < interval[0] {
			// interval 			[     ]
			// newInterval	[     ]
			res = append(res, newInterval)
			res = append(res, intervals[idx:]...)
			return res
		} else if newInterval[0] > interval[1] {
			// interval 	[     ]
			// newInterval        	[     ]
			res = append(res, interval)
		} else if newInterval[1] > interval[1] && interval[1] >= newInterval[0] && newInterval[0] >= interval[0] {
			// interval 	[     ]
			// newInterval     [     ]
			newInterval = []int{interval[0], newInterval[1]}
		} else if interval[1] >= newInterval[1] && newInterval[0] >= interval[0] {
			// interval 	[     ]
			// newInterval   [   ]
			res = append(res, intervals[idx:]...)
			return res
		} else if interval[1] >= newInterval[1] && newInterval[1] >= interval[0] && interval[0] > newInterval[0] {
			// interval 	   [     ]
			// newInterval	[     ]
			newInterval = []int{newInterval[0], interval[1]}
		} else {
			// interval 	 [   ]
			// newInterval	[     ]
		}
	}

	res = append(res, newInterval)
	return res
}

// @lc code=end

