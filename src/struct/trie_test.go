package Struct

import "testing"

func TestInsert(t *testing.T) {
	obj := NewTrie()
	msg := []string{"m", "s", "g"}
	obj.Insert(msg, "data")
}

func TestGet(t *testing.T) {
	obj := NewTrie()
	msg := []string{"m", "s", "g"}
	obj.Insert(msg, "data")
	if got, ok := obj.Get(msg); ok {
		if got.Data != "data" {
			t.Errorf("Get() = %v; want 'data'", got.Data)
		}
	}
}
