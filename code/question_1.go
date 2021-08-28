package main

import "fmt"

//@link https://leetcode-cn.com/problems/two-sum/

/**
解法1：循环计算，复杂度为O(n*n)
优化点1：比target值大的值可以不计算
 */
func twoSum(nums []int, target int) []int {
	var i = 0
	var j = 0
	var index = []int {0, 0}
	for ; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				index[0] = i
				index[1] = j
				break
			}
		}
	}
	return index
}

/**
查找
 */
func twoSumWithFind(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target - x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	var testCases = [][]int { {2, 7, 11, 15}, {3, 2, 4}, {3, 3} }
	var targets = []int {9, 6, 6}
	var i = 0
	var index = []int {0, 0}
	for ; i < len(testCases); i++ {
		index = twoSumWithFind(testCases[i], targets[i])
		fmt.Println(index)
	}
}
