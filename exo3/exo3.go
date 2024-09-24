package main

import (
	"fmt"
	"sort"
)

func Ft_non_overlap(intervalles [][]int) int {
	if len(intervalles) == 0 {
		return 0
	}

	sort.Slice(intervalles, func(i, j int) bool {
		return intervalles[i][1] < intervalles[j][1]
	})

	compteur := 0
	finDernier := intervalles[0][1]

	for i := 1; i < len(intervalles); i++ {
		if intervalles[i][0] < finDernier {
			compteur++
		} else {
			finDernier = intervalles[i][1]
		}
	}

	return compteur
}

func main() {
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // Résultat : 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // Résultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // Résultat : 2
}
