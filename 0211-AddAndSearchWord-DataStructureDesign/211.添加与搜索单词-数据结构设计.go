/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
type Trie struct {
	children [26]*Trie
	isEnd    bool
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

type WordDictionary struct {
	trieRoot *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{
		&Trie{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.trieRoot.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	var dfsSearch func(idx int, trieRoot *Trie) bool
	dfsSearch = func(idx int, trieRoot *Trie) bool {
		if trieRoot == nil {
			return false
		}

		if idx == len(word) {
			return trieRoot.isEnd
		}

		if word[idx] != '.' {
			return dfsSearch(idx+1, trieRoot.children[word[idx]-'a'])
		} else {
			for charIdx := 0; charIdx < len(trieRoot.children); charIdx += 1 {
				if dfsSearch(idx+1, trieRoot.children[charIdx]) {
					return true
				}
			}
		}
		return false
	}
	return dfsSearch(0, this.trieRoot)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

