package easy415

import "strconv"

func addStrings(num1, num2 string) string {
	add, answer := 0, ""

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0') // '0' 为 rune 类型(alias for int32)
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		answer = strconv.Itoa(result%10) + answer
		add = result / 10
	}

	return answer
}
