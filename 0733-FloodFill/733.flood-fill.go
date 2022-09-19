/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// [0 0 0]
	// [0 1 1]
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}


func dfs(image [][]int, sr int, sc int, sourceColor, newColor int) {
	if image[sr][sc] != sourceColor {
		return
	}

	if sourceColor == newColor {
		return
	}

	image[sr][sc] = newColor

    // up
	if (sr - 1) >= 0 && image[sr - 1][sc] == sourceColor {
		dfs(image, sr - 1, sc, sourceColor, newColor)
	}

	// down
	if (sr + 1) < len(image) && image[sr + 1][sc] == sourceColor {
		dfs(image, sr + 1, sc, sourceColor, newColor)
	}

	// left
	if (sc - 1) >= 0  && image[sr][sc - 1] == sourceColor {
		dfs(image, sr, sc - 1, sourceColor, newColor)
	}

	// right
	if (sc + 1) < len(image[0]) && image[sr][sc + 1] == sourceColor {
		dfs(image, sr, sc + 1, sourceColor, newColor)
	}
}
// @lc code=end

