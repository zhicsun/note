package template

import (
	"fmt"
	"reflect"
	"testing"
)

type operate struct {
	Action, Data string
}

func TestTrie(t *testing.T) {
	tests := []struct {
		name string
		ops  []operate
		want []bool
	}{
		{
			name: "测试",
			ops: []operate{
				{"insert", "apple"},
				{"search", "apple"},
				{"search", "app"},
				{"startsWith", "app"},
				{"insert", "app"},
				{"search", "app"},
			},
			want: []bool{true, false, true, true},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			trie := Trie{}
			r := make([]bool, 0)
			for _, op := range v.ops {
				fmt.Println(op)
				switch {
				case op.Action == "insert":
					trie.Insert(op.Data)
				case op.Action == "search":
					r = append(r, trie.Search(op.Data))
				case op.Action == "startsWith":
					r = append(r, trie.StartsWith(op.Data))
				}
			}
			if !reflect.DeepEqual(r, v.want) {
				t.Fatal("结果: ", r, "期望值: ", v.want)
			}
		})
	}
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	c := t
	for i := range word {
		ch := word[i] - 'a'
		if c.children[ch] == nil {
			c.children[ch] = &Trie{}
		}
		c = c.children[ch]
	}
	c.isEnd = true
}

func (t *Trie) find(word string) *Trie {
	c := t
	for i := range word {
		ch := word[i] - 'a'
		if c.children[ch] == nil {
			return nil
		}
		c = c.children[ch]
	}
	return c
}

func (t *Trie) Search(word string) bool {
	r := t.find(word)
	return r != nil && r.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != nil
}
