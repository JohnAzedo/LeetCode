// Two Sum
// https://leetcode.com/problems/two-sum/

package leetcode

// O(N^2)
func TwoSum(nums []int, target int) []int {
	var indexes []int

	for i := 0; i<len(nums); i++ {
		for j := 0; j<len(nums); j++ {
			if i == j{ continue }
			
			if (nums[i] + nums[j]) == target {
				indexes = []int{i, j}
			}
		}
	}

	return indexes
}