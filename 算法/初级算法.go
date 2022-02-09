package main

//1. 删除排序数组中的重复项
//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
//
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//func main() {
//	s := []int{1,1,2,3,3,3,4,4,4}
//	duplicates := removeDuplicates(s)
//	fmt.Println(duplicates)
//	fmt.Println(s)
//}
//
//func removeDuplicates(nums []int) int {
//	n := len(nums)
//	if n <1 {
//		return 0
//	}
//	low := 0
//	for fast := 1; fast < n; fast++ {
//		if nums[fast] != nums[low] {
//			nums[low + 1] = nums[fast]
//			low++
//		}
//		fmt.Println(nums)
//	}
//
//	return low + 1
//}
//==================================================

//2.买卖股票的最佳时机 II
//给定一个数组 prices ，其中 prices[i] 表示股票第 i 天的价格。
//
//在每一天，你可能会决定购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以购买它，然后在 同一天 出售。
//返回 你能获得的 最大 利润 。

//贪心算法
//func maxProfit(prices []int) int {
//	var sum int
//	for i := 1; i < len(prices); i++ {
//		// 累加每次大于0的交易
//		if prices[i]-prices[i-1] > 0 {
//			sum += prices[i]-prices[i-1]
//		}
//	}
//	return sum
//}
//
//func main() {
//	s := []int{7,1,5,3,6,4}
//	ans := maxProfit(s)
//	fmt.Println(ans)
//}
//==================================================

//3. 旋转数组
//给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

//func reverse(a []int) {
//	for i, n := 0, len(a); i < n/2; i++ {
//		a[i], a[n-1-i] = a[n-1-i], a[i]
//		//fmt.Println("a:",a)
//	}
//}
//func rotate(nums []int, k int) {
//	k %= len(nums)
//	reverse(nums)
//	reverse(nums[:k])
//	reverse(nums[k:])
//}
//
//func main() {
//	s := []int{1,2,3,4,5,6,7,8}
//	rotate(s,3)
//	fmt.Println(s)
//}
//==================================================

//4. 存在重复元素
//给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。

//func containsDuplicate(nums []int) bool {
//	sort.Ints(nums)
//	fmt.Println(nums)
//	for i :=1; i<len(nums); i++{
//		if nums[i] == nums[i-1]{
//			fmt.Println(nums[i], nums[i-1])
//			return true
//		}
//	}
//	return false
//}
//
//func main() {
//	s := []int{1,2,3,4,5,6,7,8,3}
//	duplicate := containsDuplicate(s)
//	fmt.Println(duplicate)
//}
//==================================================

//5. 只出现一次的数字
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//5.1 循环遍历每个数字，不是最优选择(没通过，不使用)
//func singleNumber(nums []int) int {
//	sort.Ints(nums)
//	fmt.Println(nums)
//	if len(nums) == 1{
//		return nums[0]
//	}
//	if nums[0] != nums[1]{
//		return nums[0]
//	}
//	for i := 1; i<len(nums)-1;i++{
//		if nums[i] != nums[i-1] && nums[i] != nums[i+1]{
//			return nums[i]
//		}
//	}
//	return nums[len(nums)-1]
//}
//
//func main() {
//	s := []int{-336,513,-560,-481,-174,101,-997,40,-527,-784,-283,-336,513,-560,-481,-174,101,-997,40,-527,-784,-283,354}
//	number := singleNumber(s)
//	fmt.Println(number)
//}
//5.2 使用异或运算,最优方法
//func singleNumber(nums []int) int {
//	single := 0
//	for _, num := range nums {
//		single ^= num
//		fmt.Println(single)
//	}
//	return single
//}
//==================================================
