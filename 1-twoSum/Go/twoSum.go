package twoSum

func TwoSum(nums []int, target int) []int {
	var smap = make(map[int]int)
	for i, num := range nums {
		x := target - num
		if index, ok := smap[x]; ok {
			return []int{index, i}
		}
		smap[num] = i
	}
	return []int{}
}
