package fourSum

import (
	"slices"
)


func fourSum(nums []int, target int) [][]int {
	n:= len(nums)
	if n < 4{
		return [][]int{}
	}

	slices.Sort(nums)
	results := [][]int{}


	for i := 0; i < n - 3; i++{
		if i > 0 && nums[i] == nums[i-1]{
			continue;
		}
		for j := i + 1; j < n - 2; j++{
			if j > i + 1 && nums[j] == nums[j-1]{
				continue;
			}
			left := j + 1
			right := n -1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right] 
				
				if sum == target{
					results = append(results, []int{nums[i] , nums[j] , nums[left] , nums[right] })
					
					for left < right && nums[left] == nums[left + 1]{
						left++
					}
					for left < right && nums[right] == nums[right - 1]{
						right--
					}

					left++
					right--
				} else if sum < target {
					left++
				} else{
					right--
				}


			}

		}
	}

	return results
}