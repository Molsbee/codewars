package consecutive_strings

func LongestConsec(strarr []string, k int) string {
	if len(strarr) == 0 || k <= 0 || k > len(strarr) {
		return ""
	}

	st := ""
	for i := 0; i < len(strarr); i++ {
		temp := ""
		for j := i; j < i+k; j++ {
			if j > len(strarr)-1 {
				break
			}

			temp += strarr[j]
		}

		if len(temp) > len(st) {
			st = temp
		}
	}

	return st
}
