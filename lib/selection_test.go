package kdtree

import (
	"sort"
	"testing"
)

func checkPartitionResult(t *testing.T, array sort.Interface,
	left int, right int, idx int, ans int) {
	if idx != ans {
		t.Error("Partition result index is wrong")
	}
	for i := left; i < idx; i++ {
		if !array.Less(i, idx) {
			t.Error("Partition result is wrong")
		}
	}
	for i := idx + 1; i <= right; i++ {
		if !array.Less(idx, i) {
			t.Error("Partition result is wrong")
		}
	}
}

func TestPartition(t *testing.T) {
	// case 1
	array := sort.IntSlice{2, 3, 4, 1, 5}
	idx, err := Partition(array, 0, 4, 1)
	if err != nil {
		t.Error("Partition should not return error here")
	}
	checkPartitionResult(t, array, 0, 4, idx, 2)
	// case 2
	array = sort.IntSlice{2, 3, 4, 1, 5}
	idx, err = Partition(array, 0, 4, 5)
	if err == nil {
		t.Error("Partition should return error here")
	}
	// case 3
	array = sort.IntSlice{2, 3, 4, 1, 5}
	idx, err = Partition(array, 0, 4, 4)
	if err != nil {
		t.Error("Partition should not return error here")
	}
	checkPartitionResult(t, array, 0, 4, idx, 4)
	// case 4
	array = sort.IntSlice{2, 3, 4, 1, 5}
	idx, err = Partition(array, 0, 4, 3)
	if err != nil {
		t.Error("Partition should not return error here")
	}
	checkPartitionResult(t, array, 0, 4, idx, 0)
	// case 5
	array = sort.IntSlice{2, 3, 4, 1, 5}
	idx, err = Partition(array, 2, 4, 3)
	if err != nil {
		t.Error("Partition should not return error here")
	}
	checkPartitionResult(t, array, 2, 4, idx, 2)
}

func TestQuickSelect(t *testing.T) {
	// case 1
	array := sort.IntSlice{4, 1, 2, 5, 3}
	err := QuickSelect(array, 0)
	if err == nil {
		t.Error("QuickSelect should return error here")
	}
	// case 2
	array = sort.IntSlice{4, 1, 2, 5, 3}
	err = QuickSelect(array, 6)
	if err == nil {
		t.Error("QuickSelect should return error here")
	}
	// case 3
	for i := 1; i <= 4; i++ {
		array = sort.IntSlice{4, 1, 2, 5, 3}
		err = QuickSelect(array, i)
		if err != nil {
			t.Error("QuickSelect should not return error here")
		}
		if []int(array)[i-1] != i {
			t.Error("QuickSelect result is wrong")
		}
	}
}
