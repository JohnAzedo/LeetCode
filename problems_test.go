package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T){
	result := TwoSum([]int{3, 2, 4}, 6)
	expected := []int{2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Error: %v (result) is not the same as %v (expected)", result, expected)
	} 
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	expected := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	result := AddTwoNumbers(l1, l2)
	if !reflect.DeepEqual(*result, *expected) {
		t.Fatalf("Error: %v (result) is not the same as %v (expected)", result, expected)
	}
}