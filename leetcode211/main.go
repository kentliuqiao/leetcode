package main

func main() {}

type WordDictionary struct {
	isWord   bool
	children [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		isWord:   false,
		children: [26]*WordDictionary{},
	}
}

/** adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}
	for _, ch := range word {
		if this.children[ch-'a'] == nil {
			this.children[ch-'a'] = new(WordDictionary)
		}
		this = this.children[ch-'a']
	}
	this.isWord = true
}

/** Returns if the word is in the data structure. a word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return true
	}

	return this.match(word, 0)
}

func (this *WordDictionary) match(word string, start int) bool {
	if len(word) == start {
		return this.isWord
	}

	ch := word[start]
	if ch == '.' {
		for i := 0; i < 26; i++ {
			if this.children[i] == nil {
				continue
			}
			if this.children[i].match(word, start+1) {
				return true
			}
		}
		return false
	}
	this = this.children[ch-'a']
	if this == nil {
		return false
	}

	return this.match(word, start+1)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.addWord(word);
 * param_2 := obj.Search(word);
 */
