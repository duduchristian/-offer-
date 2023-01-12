package main

import (
	"algorithm/offer"
	"fmt"
	"math"
)

func main() {
	fmt.Println(offer.AddBinary("1001", "1010"))
	//fmt.Println(offer.SingleNumber([]int{-1, -1, -1, 0, 0, 0, math.MaxInt}))
	//fmt.Println(offer.MaxProduct([]string{"abcw", "foo", "bar", "fxyz", "abcdef"}))
	//fmt.Println(offer.SingleNumberBruteForce([]int{-1, -1, -1, 0, 0, 0, math.MaxInt}))
	//fmt.Println(offer.TwoSum([]int{1, 2, 4, 6, 10}, 8))
	//fmt.Println(offer.ThreeSum([]int{-1, 0, 1, 2, -1, -4, 0, 0}))
	//fmt.Println(offer.MinSubArrayLen(7, []int{5, 1, 4, 3}))
	//fmt.Println(offer.MinSubArrayLen(7, []int{9, 7, 9, 7}))
	//fmt.Println(offer.SubarraySum([]int{1, 1, 1}, 2))
	//fmt.Println(offer.SubarraySum([]int{0, 0, 1, 0}, 0))
	//fmt.Println(offer.FindMaxLength([]int{0, 1, 0, 1, 0, 1, 0, 1, 1, 0}))
	//fmt.Println(offer.LengthOfLongestSubstring("abcdefg"))
	//fmt.Println(offer.MinWindow("ADDBANCAD", "ABCD"))
	fmt.Println(offer.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(offer.DailyTemperatures([]int{35, 31, 33, 36, 34}))
	fmt.Println(math.Pow(1.3, 104))
	fmt.Println(offer.KSmallestPairs([]int{1, 2}, []int{3, 4}, 20))
}
