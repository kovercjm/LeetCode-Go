package main

func peakIndexInMountainArray(arr []int) int {
	return binarySearch(arr, 0, len(arr))
}

func binarySearch(arr []int, i, j int) int {
	if j-i < 2 {
		return -1
	}
	mid := (i + j) / 2
	if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
		return mid
	} else if arr[mid-1] > arr[mid] {
		return binarySearch(arr, i, mid)
	} else {
		return binarySearch(arr, mid, j)
	}
}
