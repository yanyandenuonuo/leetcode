/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	// solution1: 模拟
	// [availableTime, remainedCounter]
	/**
	taskMap := make(map[byte]bool)
	taskList := [26][2]int{}
	for _, task := range tasks {
		taskList[task-'A'][0] = 1
		taskList[task-'A'][1] += 1

		taskMap[task] = true
	}

	currentTime := 0
	for len(taskMap) > 0 {
		// 选取可执行的剩余任务最多的一个任务
		minAvailableTime, taskName, remainedCounter := len(tasks)*(n+1), byte(0), 0
		for currentTaskName := range taskMap {
			if taskList[currentTaskName-'A'][0] < minAvailableTime {
				minAvailableTime = taskList[currentTaskName-'A'][0]
			}

			if taskList[currentTaskName-'A'][0] <= currentTime && taskList[currentTaskName-'A'][1] > remainedCounter {
				taskName = currentTaskName
				remainedCounter = taskList[currentTaskName-'A'][1]
			}
		}

		if minAvailableTime > currentTime {
			// 不存在可以执行的task
			currentTime = minAvailableTime
			continue
		}

		taskList[taskName-'A'][1] -= 1
		if taskList[taskName-'A'][1] == 0 {
			delete(taskMap, taskName)
		} else {
			taskList[taskName-'A'][0] = currentTime + 1 + n
		}

		currentTime += 1
	}

	return currentTime - 1
	*/

	// solution2: 构造，按列构造任务，最小时间为len(tasks)和(maxExecCounter-1)*(n+1)+sameMaxExecCounter的较大值
	//			  A B C
	//			  A B C
	//			  A B C
	//			  A B C
	maxExecCounter, sameMaxExecCounter := 0, 0
	taskCounterList := [26]int{}
	for _, task := range tasks {
		taskCounterList[task-'A'] += 1

		if taskCounterList[task-'A'] > maxExecCounter {
			maxExecCounter = taskCounterList[task-'A']
			sameMaxExecCounter = 1
		} else if taskCounterList[task-'A'] == maxExecCounter {
			sameMaxExecCounter += 1
		}
	}

	taskCounter := (maxExecCounter-1)*(n+1) + sameMaxExecCounter
	if len(tasks) > taskCounter {
		return len(tasks)
	}

	return taskCounter
}

// @lc code=end

