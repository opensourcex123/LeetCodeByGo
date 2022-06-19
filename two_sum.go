package main

import "log"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		another := target - num
		if _, ok := m[another]; ok {
			return []int{m[another], index}
		}
		m[nums[index]] = index
	}

	return nil
}

func main() {
	num := twoSum([]int{1, 2, 3}, 4)
	log.Println(num)
}
