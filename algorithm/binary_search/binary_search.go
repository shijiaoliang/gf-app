package binary_search

import (
	"fmt"
)

func binarySearch(slice []int, target int, l, r int) int {
	if l > r {
		return -1
	}

	mid := (l + r) / 2
	midValue := slice[mid]

	if target == midValue {
		return mid
	} else if midValue > target {
		return binarySearch(slice, target, l, mid-1)
	} else {
		return binarySearch(slice, target, mid+1, r)
	}
}

func binarySearch2(slice []int, target int, l, r int) int {
	for {
		if l > r {
			return -1
		}

		mid := (l + r) / 2
		midValue := slice[mid]

		if midValue == target {
			return mid
		} else if midValue > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
}

func main() {
	a := []int{10, 20, 30, 31, 45}
	fmt.Println(binarySearch2(a, 31, 0, len(a)-1))
}
