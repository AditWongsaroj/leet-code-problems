package threeSumClosest

import (
	"math"
	"sort"
)


func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closest := nums[0] + nums[1] + nums[2]
	difference := math.Abs(float64(target-closest))

	for i := range nums{
		left := i+1
		right := len(nums)-1

		t := target - nums[i]

		for left < right{
			sum := nums[left] + nums[right]

			if sum == t {return target}
			if sum < t { left++}else {right--} 

			diff := math.Abs(float64(t-sum))
			
			if(diff < difference){
				closest = sum + nums[i]
				difference = diff
			}
		}
	}
	return int(closest)
}

// Brute Force

// func threeSumClosest(nums []int, target int) int {
// 	closest := nums[0]+nums[1]+nums[2]
// 	distance := math.Abs(float64(closest-target))
// 	for i, _ := range nums{
// 		for j, _ := range nums{
// 			for k, _ := range nums{
// 				if k > j && j > i {
// 					current := nums[i]+nums[j]+nums[k]
// 					cur_dist := math.Abs(float64(current-target))
// 					if cur_dist < distance {
// 						closest = current
// 						distance = cur_dist
// 					}
// 				}
// 	}
// 	}
// 	}
// 	return int(closest)
// }