package medium43

import "strconv"

func multiply(num1, num2 string) string {
	answer := `0`
	if num1 == `0` || num2 == `0` {
		return answer
	}
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- { // num2[i]
		result, carry := ``, 0
		// 通过初始尾部补零来规避成 10 的倍数操作，可规避溢出隐患
		for j := i + 1; j <= n-1; j++ {
			result += `0`
		}
		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- { // num1[j]
			x := int(num1[j] - '0')
			product := x*y + carry
			carry = product / 10
			result = strconv.Itoa(product%10) + result
		}

		for ; carry != 0; carry /= 10 {
			result = strconv.Itoa(carry%10) + result
		}

		answer = addStrings(answer, result)
	}

	return answer
}

func addStrings(x, y string) string {
	carry, answer := 0, ``

	for i, j := len(x)-1, len(y)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		a, b := 0, 0
		if i >= 0 {
			a = int(x[i] - '0')
		}
		if j >= 0 {
			b = int(y[j] - '0')
		}
		sum := a + b + carry
		carry = sum / 10
		answer = strconv.Itoa(sum%10) + answer
	}

	return answer
}
