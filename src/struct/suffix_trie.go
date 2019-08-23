package Struct

type SuffixTrie struct {
	T *Trie
}

func NewSuffixTrie() *SuffixTrie {
	return &SuffixTrie{
		T: NewTrie(),
	}
}

func (t *SuffixTrie) Insert(s []string, data interface{}) {
	t.T.Insert(Reverse(s), data)
}

func (t SuffixTrie) Get(s []string) (*Node, bool) {
	return t.T.Get(Reverse(s))
}

func Reverse(l []string) []string {
	for i := len(l)/2 - 1; i >= 0; i-- {
		opp := len(l) - 1 - i
		l[i], l[opp] = l[opp], l[i]
	}
	return l
}
