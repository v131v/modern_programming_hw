package bubblesort

func Sort(ar []int, cmp func(i int, j int) bool) {

	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if !cmp(i, j) {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
}
