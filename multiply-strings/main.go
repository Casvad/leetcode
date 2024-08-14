package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ansArr := make([]int, len(num1)+len(num2))

	for i, m := range num1 {
		for j, n := range num2 {
			arrIndex := i + j
			mstr := string(m)
			mint, _ := strconv.Atoi(mstr)
			nstr := string(n)
			nint, _ := strconv.Atoi(nstr)
			multi := mint * nint
			if multi > 10 {
				if arrIndex == len(ansArr)-1 {
					ansArr[arrIndex] += multi
				} else {
					ansArr[arrIndex+1] += multi % 10
					ansArr[arrIndex] += multi / 10
				}
			} else {
				ansArr[arrIndex+1] += multi
			}
		}
	}

	ans := ""
	acuumulative := 0
	fmt.Printf("%v\n", ansArr)
	for i := len(ansArr) - 1; i >= 0; i-- {
		acuumulative += ansArr[i]
		if acuumulative == 0 && i == 0 {
			break
		}
		ans = fmt.Sprintf("%v", int(acuumulative%10)) + ans
		acuumulative = acuumulative / 10
	}
	if acuumulative > 0 {
		return fmt.Sprintf("%v", acuumulative) + ans
	}

	return ans
}

func main() {

}
