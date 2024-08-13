package main

import "fmt"

func bubleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func QvickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	p := nums[0]

	l := []int{}
	g := []int{}
	c := []int{}

	for _, v := range nums {
		if v < p {
			l = append(l, v)
		} else if v > p {
			g = append(g, v)
		} else if v == p {
			c = append(c, v)
		}
	}

	return append(append(QvickSort(l), c...), QvickSort(g)...)

}

func deleteDuble(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
			i--
		}
	}

	return nums
}

func search(nums []int, t int) int {
	a := 0
	b := len(nums) - 1

	for a <= b {
		mid := (a + b) / 2
		if nums[mid] == t {
			return mid
		} else if nums[mid] > t {
			b = mid - 1
		} else if nums[mid] < t {
			a = mid + 1
		}
	}

	return -1 // not found
}

func main() {

	nums := []int{1, 20, 2, 4, 10, 4, 12}
	t := 4

	sortedNums := bubleSort(nums)
	sortedNumsQvick := QvickSort(nums)
	setNums := deleteDuble(sortedNums)
	seartchElement := search(setNums, t)

	fmt.Println("Index:", seartchElement)
	fmt.Println(sortedNumsQvick)
	fmt.Println(sortedNums)

}
