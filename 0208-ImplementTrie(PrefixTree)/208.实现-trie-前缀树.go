/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	currentPoint := this
	for _, runeChar := range word {
		if currentPoint.children[runeChar-'a'] == nil {
			currentPoint.children[runeChar-'a'] = &Trie{}
		}
		currentPoint = currentPoint.children[runeChar-'a']
	}
	currentPoint.isEnd = true
}

func (this *Trie) Search(word string) bool {
	currentPoint := this
	for _, runeChar := range word {
		if currentPoint == nil {
			return false
		}
		currentPoint = currentPoint.children[runeChar-'a']
	}
	return currentPoint != nil && currentPoint.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	currentPoint := this
	for _, runeChar := range prefix {
		if currentPoint == nil {
			return false
		}
		currentPoint = currentPoint.children[runeChar-'a']
	}
	return currentPoint != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

