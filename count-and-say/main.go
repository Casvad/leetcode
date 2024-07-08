package main

import "fmt"

func countAndSay(n int) string {

	ans := rle("")
	for i := 1; i < n; i++ {
		ans = rle(ans)
	}

	return ans
}

func rle(str string) string {

	if str == "" {
		return "1"
	}

	ansStr := ""
	i := 1
	r := str[0]
	currentCount := 1
	for i < len(str) {
		if r == str[i] {
			currentCount++
		} else {
			ansStr += fmt.Sprintf("%d%s", currentCount, string(r))
			currentCount = 1
			r = str[i]
		}
		i++
	}
	ansStr += fmt.Sprintf("%d%s", currentCount, string(r))

	return ansStr
}

func main() {

	fmt.Printf("4: %s\n", countAndSay(4))
}
