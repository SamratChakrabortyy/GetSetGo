package main

func groupAnagrams(strs []string) [][]string {

	m := make(map[[26]int][]string)

	for _, str := range strs {
		hash := [26]int{}
		for _, s := range str {
			hash[s-'a']++
		}
		m[hash] = append(m[hash], str)
	}
	ans := [][]string{}
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}
