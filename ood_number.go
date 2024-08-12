package main

func findCountOdd(listNum []int) (max int) {
	isOdd := func(n int) bool {
		return n%2 != 0
	}

	mapNum := make(map[int]int)

	for _, v := range listNum {
		if val, ok := mapNum[v]; !ok {
			mapNum[v] = mapNum[v] + 1
		} else {
			mapNum[v] = val + 1
		}
	}

	for key, val := range mapNum {
		if isOdd(val) {
			max = key
		}
	}

	return

}
