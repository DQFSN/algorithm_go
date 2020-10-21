func permutation(s string) []string {

	chars := []rune(s)

	result := make([][]rune, 0, 20)
	resultString := make([]string, 0, 20)
	mapStr := make(map[string]bool)

	for i := 0; i < len(chars); i++ {

		newResults := make([][]rune, 0)

		for _, r := range result {
			//剪枝
			if len(r) != i {
				continue
			}

			r = append(r, chars[i])

			for j, _ := range r {
				//深拷贝
				cp := make([]rune, len(r))
				_ = copy(cp, r)
				cp[j], cp[len(cp)-1] = cp[len(cp)-1], cp[j]
				newResults = append(newResults, cp)
				// fmt.Println(string(r), "---", j, "----", string(cp))
			}
		}

		if len(result) == 0 {
			first := []rune{chars[0]}
			result = append(result, first)
			if len(first) == len(s) {
				resultString = append(resultString, string(first))
			}
		}

		for _, nr := range newResults {
			result = append(result, nr)
			strNr := string(nr)
			if _, ok := mapStr[strNr]; !ok && len(nr) == len(s) {
				resultString = append(resultString, strNr)
				mapStr[strNr] = true
			}
		}

	}

	if len(resultString) == 0 {
		return []string{""}
	}

	return resultString
}
