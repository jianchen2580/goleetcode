package 00001_two_sum

func twoSum(nums []int, target int) []int {
	for i, key := range nums {

		for j, value := range nums[i+1:] {
			if target-key == value {
				return []int{i, j + i + 1}
			}
		}
	}
  return nil
}

