package qsort

import "testing"

func TestQuickSort1(t *testing.T)  {
	values := [] int {5,4,3,2,1}
	QuickSort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5{
		t.Error("quick Sort failed, got ", values, "Expect {1,2,3,4,5}")
	}

}

func TestQuickSort2(t *testing.T)  {
	values := [] int {5,5,4,3,2,1}
	QuickSort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 || values[5] != 5{
		t.Error("quick Sort failed, got ", values, "Expect {1,2,3,4,5, 5}")
	}

}

func TestQuickSort3(t *testing.T)  {
	values := [] int {5}
	QuickSort(values)

	if values[0] != 5{
		t.Error("quick Sort failed, got ", values, "Expect {5}")
	}

}

func TestQuickSort4(t *testing.T)  {
	values := make([] int, 0)
	QuickSort(values)

	if len(values) != 0{
		t.Error("quick Sort failed, got ", values, "Expect {}")
	}

}
