package levdist

func Calculate(a string, b string) int {
	return levDistTwoRows(a, b)
}

func levDistSlice(a string, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	var indicator int
	if a[0] != b[0] {
		indicator = 1
	} else {
		indicator = 0
	}

	return min(
		levDistSlice(a[1:], b)+1,
		levDistSlice(a, b[1:])+1,
		levDistSlice(a[1:], b[1:])+indicator,
	)
}

func levDistMatrix(a string, b string) int {
	n, m := len(a)+1, len(b)+1
	distance := make([]int, n*m, n*m)
	for i := 1; i < n; i++ {
		distance[i] = i
	}
	for j := 1; j < m; j++ {
		distance[j*n] = j
	}

	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			var indicator int
			if a[i-1] != b[j-1] {
				indicator = 1
			} else {
				indicator = 0
			}

			distance[i+j*n] = min(
				distance[(i-1)+j*n]+1,
				distance[i+(j-1)*n]+1,
				distance[(i-1)+(j-1)*n]+indicator,
			)
		}
	}

	return distance[n*m-1]
}

func levDistTwoRows(a string, b string) int {
	n := len(a)
	m := len(b)
	previous := make([]int, m+1, m+1)
	for i := 0; i <= m; i++ {
		previous[i] = i
	}
	for i := range n {
		current := make([]int, m+1, m+1)
		current[0] = i + 1
		for j := 1; j <= m; j++ {
			var indicator int
			if a[i] != b[j-1] {
				indicator = 1
			} else {
				indicator = 0
			}

			current[j] = min(
				previous[j]+1,
				current[j-1]+1,
				previous[j-1]+indicator,
			)
		}
		previous = current
	}

	return previous[m]
}

func min(a int, b int, c int) int {
	smallest := a
	if b < smallest {
		smallest = b
	}
	if c < smallest {
		smallest = c
	}

	return smallest
}
