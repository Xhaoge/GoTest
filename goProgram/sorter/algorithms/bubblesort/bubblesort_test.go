package bubblesort

import (
	"testing"
)

func TestBubbleSort1(t * testing.T) {
	values := []int{5,4,3,2,1,8}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3{
		t.Error("BubbleSort failed ",values,"expected 1,2,3,4,5,8")
	}
} 

func TestBubbleSort2(t *testing.T){
	values := []int{5,5,3,2,1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 {
		t.Error("BubbleSort failed ",values,"expected 1,2,3,5,5")
	}
}

func TestBubbleSort3(t *testing.T){
	values := []int{5}
	BubbleSort(values)
	if values[0] != 5{
		t.Error("BubbleSort failed ",values,"expected 5")
	}
}