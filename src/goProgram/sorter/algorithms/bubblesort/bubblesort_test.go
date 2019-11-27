package bubblesort

import (
	"testing"
)

func TestBubbleSort1(t * testing.T) {
	values := []int{5,4,3,2,1,8}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3{
		
	}
} 