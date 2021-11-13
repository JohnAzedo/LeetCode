package datastructure

import (
	"reflect"
	"testing"
)

func TestLinkedListAddFirst(t *testing.T) {
	list := New()
	list.AddFirst(1)
	expected := &Node{1, nil}
	if !reflect.DeepEqual(*list.first, *expected){
		t.Fatalf("Error: %v (result) is not the same as %v (expected)", list.first, expected)
	}
}

func TestLinkedListAddLast(t *testing.T){
	list := New()
	list.AddFirst(1)
	list.AddLast(2)
	expected := &Node{1, &Node{2, nil}}
	if !reflect.DeepEqual(*list.first, *expected){
		t.Fatalf("Error: %v (result) is not the same as %v (expected)", list.first, expected)
	}
}

func TestLinkedListAddLastWithEmptyList(t *testing.T){
	list := New()
	list.AddLast(2)
	expected := &Node{2, nil}
	if !reflect.DeepEqual(*list.first, *expected){
		t.Fatalf("Error: %v (result) is not the same as %v (expected)", list.first, expected)
	}
}