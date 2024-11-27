package main

func sort74(x []int) []int {
	k := 1
	n := len(x)

	for k < n {
		t := 0

		for t+k < n {

			p := t
			q := t + k - 1
			r := t + 2*k - 1

			if r >= n {
				r = n - 1
			}
			merge(x, p, q, r)
			t = r + 1
		}
		k *= 2
	}

	return x
}

func merge(x []int, p, q, r int) {
	left := make([]int, q-p+1)
	right := make([]int, r-q)

	copy(left, x[p:q+1])
	copy(right, x[q+1:r+1])

	i, j, k := 0, 0, p

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			x[k] = left[i]
			i++
		} else {
			x[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		x[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		x[k] = right[j]
		j++
		k++
	}
}
