package easy941

func validMuntainArray(A []int) bool {
	i, n := 0, len(A)

	for i+1 < n && A[i] < A[i+1] {
		i++
	}

	if i == 0 || i == n-1 {
		return false
	}

	for i+1 < n && A[i] > A[i+1] {
		i++
	}

	return i == n-1
}
