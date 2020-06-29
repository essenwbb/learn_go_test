package leetcode

import (
	"container/heap"
	"fmt"
)

func letterCombinations(digits string) []string {
	dict := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	var result []string
	for _, value := range digits {
		if result == nil {
			for _, tmp := range dict[string(value)] {
				result = append(result, tmp)
			}
		} else {
			var newValue []string
			for _, res := range result {
				for _, tmp := range dict[string(value)] {
					newValue = append(newValue, res+tmp)
				}
			}
			result = append([]string{}, newValue...)
		}
	}
	fmt.Print('3' - '2')
	return result
}

type TopList []int

func (t TopList) Len() int {
	return len(t)
}

func (t TopList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TopList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t *TopList) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func (t *TopList) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func findKthLargest(nums []int, k int) int {
	m := make(TopList, 0)
	size := 0
	for i := range nums {
		if size < k {
			heap.Push(&m, nums[i])
			size++
		} else {
			if nums[i] > m[0] { //小顶堆 堆顶元素小于当前元素
				heap.Pop(&m)
				heap.Push(&m, nums[i])
			}
		}
	}
	return m[0]
}
