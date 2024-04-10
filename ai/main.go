package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	charIndex := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < n; right++ {
		if index, found := charIndex[s[right]]; found && index >= left {
			left = index + 1
		}
		charIndex[s[right]] = right
		if length := right - left + 1; length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func getMaxColors(prices []int32, money int32) int32 {

	ans := int32(0)

	for i := 0; i < len(prices); i++ {

		for j := i; j < len(prices); j++ {
			obtainedMoney := int32(0)
			ans2 := int32(0)
			k := j
			for {
				if k >= len(prices) || obtainedMoney+prices[k] > money {
					break
				}
				ans2++
				obtainedMoney += prices[k]
				k++
			}
			if ans2 > ans {
				ans = ans2
			}
		}
	}

	return ans
}

func main() {

	ans := getMaxColors([]int32{
		10, 10, 10,
	}, 5)

	ans2 := getMaxColors([]int32{
		33, 42, 21, 21, 45, 49, 1, 1, 4, 13, 45, 35, 39, 33, 33, 16, 20, 50, 13, 42, 2, 31, 44, 17, 28, 39, 4, 2, 22, 29, 42, 4, 21, 12, 26, 15, 10, 27, 17, 13, 41, 13, 50, 29, 45, 34, 46, 16, 33, 10, 7, 34, 41, 3, 2, 18, 41, 6, 21, 14, 34, 12, 19, 6, 25, 45, 20, 37, 21, 39, 49, 13, 3, 50, 43, 48, 33, 38, 15, 17, 47, 22, 3, 37, 24, 4, 6, 16, 11, 26, 30, 45, 39, 48, 50, 14, 44, 22, 50, 16,
	}, 50)

	print(ans)
	print(ans2)
}
