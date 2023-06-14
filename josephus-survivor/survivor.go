package josephus_survivor

func JosephusSurvivor(n, k int) int {
	people := make([]int, n)
	for i := range people {
		people[i] = i + 1
	}

	position := -1
	for {
		if len(people) == 1 {
			break
		}

		for i := 0; i < k; i++ {
			position++
			if position == len(people) {
				position = 0
			}
		}

		people = append(people[:position], people[position+1:]...)
		position -= 1
	}
	return people[0]
}
