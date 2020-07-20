package trie

// Node 前缀树/字典树/trie树
type Node struct {
	children [26]*Node
	isEnd    bool
}

// Constructor Initialize your data structure here
func Constructor() Node {
	return Node{}
}

// Insert ...
func (t *Node) Insert(word string) {
	for _, ch := range word {

		if t.children[ch-'a'] == nil {
			t.children[ch-'a'] = new(Node)
		}
		t = t.children[ch-'a']
	}
	t.isEnd = true
}

// Search ...
func (t *Node) Search(word string) bool {
	for _, ch := range word {
		if t.children[ch-'a'] == nil {
			return false
		}
		t = t.children[ch-'a']
	}

	return t.isEnd
}

// StartsWith ...
func (t *Node) StartsWith(prefix string) bool {
	for _, ch := range prefix {
		if t.children[ch-'a'] == nil {
			return false
		}
		t = t.children[ch-'a']
	}
	return true
}
