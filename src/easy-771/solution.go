package easy771

func numJewelsInStones(jewels, stones string) (sum int) {
	kinds := make(map[rune]int)
	for _, jewel := range jewels {
		kinds[jewel]++
	}
	for _, stone := range stones {
		if _, ok := kinds[stone]; ok {
			sum++
		}
	}
	return
}
