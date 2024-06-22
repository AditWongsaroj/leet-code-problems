package threeSum

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	results := [][]int{}
	start := nums[0]
	if start == 0 && nums[1] == 0 && nums[2] == 0 && len(nums) >= 3 { return [][]int{{0,0,0}}}
	if start >= 0 {return results}

	last_num := len(nums) - 1
	if nums[last_num] == 0 && nums[last_num-1] == 0 && nums[last_num-2] == 0 && len(nums) >= 3 { return [][]int{{0,0,0}}}
	if nums[last_num] <= 0 {return results}
	
	start = -5000 //reset start to setup duplicate check
	for i := range nums {
		if start == nums[i]{
			continue
		}
		start = nums[i]
		if start == 0 {
			if start == 0 && nums[i + 1] == 0 && nums[i + 2] == 0{
				results = append(results, []int{0,0,0})
			}
			break
		}
		if start > 0 {break}
		left := i + 1
		right := last_num
		
		for left < right {
			// mid := int(right/2) - int(left/2);	
			
			r_num := nums[right] 
			l_num := nums[left] 

			sum := start + l_num +r_num
			
			if sum == 0 {
				results = append(results, []int{start, l_num, r_num})
			}	

			if sum < 0 {
				left++
                for nums[left] == l_num && left < right{
                    left++
                }
			} else {
				right--               
                for nums[right] == r_num && left < right{
                    right--
                }
			}
		}
	}
	return results
}

