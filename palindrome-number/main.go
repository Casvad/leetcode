package main

import (
	"fmt"
)

func main() {

	runExample("a", ".*..a*")                                  //true
	runExample("aaaaaaaaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*") //true
	runExample("a", "ab*")                                     //true
	runExample("aaa", "a.a")                                   //true
	runExample("ab", ".*c")                                    //false
	runExample("mississippi", "mis*is*p*.")                    //false
	runExample("aa", "a")                                      //false
	runExample("aa", "a*")                                     //true
	runExample("ab", ".*")                                     //true
	runExample("aab", "c*a*b")                                 //true
}

func runExample(s string, p string) {
	result := isMatch(s, p)
	fmt.Printf("The result of %v %v is %v \n", s, p, result)
}

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j] || dp[i][j]
				}
			}
		}
	}

	return dp[m][n]
}

func isMatchV1(s string, p string) bool {

	if len(s) == 0 {
		index := 0
		for index < len(p) {
			if index+1 < len(p) && string(p[index+1]) == "*" {
				index += 2
			} else {
				return false
			}
		}

		return true
	}
	if len(p) <= 1 {
		return len(s) == 1 && (s == p || p == ".")
	}

	if string(p[1]) == "*" {
		if len(s) == 0 {
			return isMatch(s, p[2:])
		}
		index := 0
		canContinue := true
		for index < len(s) && canContinue {
			match := false
			if string(p[0]) == "." || string(p[0]) == string(s[index]) {
				match = isMatch(s[index+1:], p[2:]) || isMatch(s[index:], p[2:])
			} else {
				canContinue = false
				match = isMatch(s[:], p[2:])
			}
			if match {
				return true
			}
			index++
		}
		return false
	}

	return (string(p[0]) == string(s[0]) || string(p[0]) == ".") && isMatch(s[1:], p[1:])
}
