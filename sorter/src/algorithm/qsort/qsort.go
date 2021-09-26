package qsort
func quickSort(values [] int, left int, right int){
	if values == nil ||len(values) == 0{
		return
	}
	temp := values[left]
	p := left
	i, j := left, right
	for i <= j{
		for j>=p && values[j] > temp{
			j --
		}
		if j >= p{
			values[p] = values[j]
			p = j
		}
		if values[j] <= temp && i <= p{
			i ++
		}
		if i <= p{
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p - left > 1{
		quickSort(values, left, p-1)
	}
	if right - p > 1{
		quickSort(values, p + 1, right)
	}
}

func QuickSort(values [] int)  {
	quickSort(values, 0, len(values) - 1)
}