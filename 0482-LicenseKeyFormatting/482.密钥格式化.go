/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(s string, k int) string {
	strBuild := new(strings.Builder)

	for _, runeChar := range s {
		if runeChar == '-' {
			continue
		}

		if runeChar >= 'a' && runeChar <= 'z' {
			runeChar = 'A' + runeChar - 'a'
		}

		strBuild.WriteByte(byte(runeChar))
	}
	s = strBuild.String()

	keyList := make([]string, 0, len(s)/k+1)
	if len(s)%k > 0 {
		keyList = append(keyList, s[:len(s)%k])
	}

	for idx := len(s) % k; idx < len(s); idx += k {
		keyList = append(keyList, s[idx:idx+k])
	}

	return strings.Join(keyList, "-")
}

// @lc code=end

