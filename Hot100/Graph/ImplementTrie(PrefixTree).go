package Graph

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isEnd bool
	next  [26]*TrieNode
}

func initTrieNode() *TrieNode {
	return &TrieNode{
		isEnd: false,
		next:  [26]*TrieNode{},
	}
}

func Constructor() Trie {
	return Trie{
		root: initTrieNode(),
	}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	cur := this.root
	for i := 0; i < n; i++ {
		if cur.next[word[i]-'a'] == nil {
			cur.next[word[i]-'a'] = initTrieNode()
		}
		cur = cur.next[word[i]-'a']
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	n := len(word)
	cur := this.root
	for i := 0; i < n; i++ {
		if cur.next[word[i]-'a'] == nil {
			return false
		}
		cur = cur.next[word[i]-'a']
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	n := len(prefix)
	cur := this.root
	for i := 0; i < n; i++ {
		if cur.next[prefix[i]-'a'] == nil {
			return false
		}
		cur = cur.next[prefix[i]-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
