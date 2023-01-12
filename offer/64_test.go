package offer

import "testing"

func TestMagicDictionary_Search(t *testing.T) {
	dict := NewMagicDictionary([]string{"happy", "new", "year"})
	if dict.Search("happy") {
		t.Fatal("Error!")
	}
	if !dict.Search("now") {
		t.Fatal("Error!!")
	}
	if !dict.Search("yeah") {
		t.Fatal("Error!!!")
	}
	if dict.Search("hello") {
		t.Fatal("Error!!!!")
	}
}
