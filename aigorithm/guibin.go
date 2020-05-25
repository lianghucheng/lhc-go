package main

import "fmt"

func main() {
	testData := []int{6,5,4,3,8,3,8,1,2,3,4,11,17}
	mergeSort(testData, 0,len(testData)-1)
	fmt.Println(testData)
}

func mergeSortErr(arr []int, lo, mid, hi int) {
	if lo + 1 < hi {//为什么是错误示例，因为我这样子写，当出现数组长度为1的情况时，会继续进行合并操作，而此时进行合并操作会导致越界。正确做法是，判断完边界条件之后，不进行合并操作
		mergeSortErr(arr, lo, (lo+mid)/2,mid)
		mergeSortErr(arr, mid +1,(mid + hi +1)/2, hi)
	}

	i:= lo
	j := mid+1
	k := lo
	arr1 := arr[lo:mid +1]
	arr2 := arr[mid +1:hi]

	for {
		if i > mid {
			arr[k] = arr2[j]
			k++
			j++
		} else if j > hi {
			arr[k] = arr1[i]
			k++
			i++
		}else if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			k++
			i++
		} else if arr1[i] > arr2[j] {
			arr[k] = arr2[j]
			k++
			j++
		}
	}
}

//出现的问题：1、边界条件搞错。2、合并未考虑相等情况
func mergeSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (hi + lo)/2
	mergeSort(arr, lo, mid)
	mergeSort(arr, mid +1, hi)

	merge(arr, lo, mid, hi)
}

func merge(arr []int, lo, mid, hi int){
	remain := append([]int{}, arr...)
	k,i,j := lo,lo,mid +1
	for k <= hi {
		if i > mid {
			arr[k] = remain[j]
			k,j = k+1,j+1
		} else if j > hi {
			arr[k] = remain[i]
			k, i = k + 1, i + 1
		} else if remain[i] < remain[j] {
			arr[k] = remain[i]
			k, i = k+1,i+1
		} else {
			arr[k] = remain[j]
			k, j = k+1, j+1
		}
	}
}