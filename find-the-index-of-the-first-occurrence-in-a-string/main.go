package main

func main() {

}

func strStr(haystack string, needle string) int {

	for i := range haystack {

		j := 0

		for {
			if i+j >= len(haystack) || j >= len(needle) || haystack[i+j] != needle[j] {
				break
			}
			j++
		}
		if j == len(needle) {
			return i
		}
	}

	return -1
}
