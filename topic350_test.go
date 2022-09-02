package luoqiangLikou

import "testing"

func TestTopic350(T *testing.T) {

}

func intersect(nums1 []int, nums2 []int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		count[nums1[i]]++
	}

	var res []int
	for i := 0; i < len(nums2); i++ {
		if count[nums2[i]] > 0 {
			count[nums2[i]]--
			res = append(res, nums2[i])
		}
	}
	return res
}
