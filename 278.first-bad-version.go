/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    rightIdx := n
	for leftIdx := 1; leftIdx <= rightIdx; {
		midIdx := leftIdx + (rightIdx - leftIdx) / 2

		if isBadVersion(midIdx) {
			if midIdx == 1 || !isBadVersion(midIdx - 1){
				return midIdx
			} else {
				rightIdx = midIdx - 1
			}
		} else {
			leftIdx = midIdx + 1
		}
	}
	return -1
}
// @lc code=end

