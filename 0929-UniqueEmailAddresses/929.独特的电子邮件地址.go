/*
 * @lc app=leetcode.cn id=929 lang=golang
 *
 * [929] 独特的电子邮件地址
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	emailMap := make(map[string]bool)

	for _, email := range emails {
		realEamil := getRealEmail(email)
		if realEamil == "" {
			continue
		}

		emailMap[realEamil] = true
	}

	return len(emailMap)
}

func getRealEmail(email string) string {
	realEmailBuilder := new(strings.Builder)
	ignoreChar, ignoreDot, atCount := false, true, 0
	for _, runeChar := range email {
		switch runeChar {
		case '.':
			if ignoreDot {
				continue
			}
		case '+':
			if atCount == 0 {
				ignoreChar = true
				continue
			}
		case '@':
			atCount += 1
			ignoreChar = false
			ignoreDot = false
		default:
			if atCount == 0 && ignoreChar {
				continue
			}
		}

		realEmailBuilder.WriteRune(runeChar)
	}

	if atCount > 1 {
		return ""
	}

	realEmail := realEmailBuilder.String()
	if realEmail[0] == '@' {
		return ""
	}

	return realEmail
}

// @lc code=end

