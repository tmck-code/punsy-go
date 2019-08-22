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
		if got.Data[0] != "data" {
			t.Errorf("Get() = %v; want 'data'", got.Data)
		}
	}
}

func TestInsertWithSameWord(t *testing.T) {
	obj := NewTrie()
	msg := []string{"a", "b"}
	obj.Insert(msg, "cat")
	obj.Insert(msg, "dog")
	want := [2]string{"cat", "dog"}
	if got, ok := obj.Get(msg); ok {
		for _, exp := range want {
			if !Contains(got.Data, exp) {
				t.Errorf("Get() = %v; want %v", got.Data, want)
			}
		}
	}
}

func Contains(slice []interface{}, s string) bool {
	for _, el := range slice {
		if el == s {
			return true
		}
	}
	return false
}
