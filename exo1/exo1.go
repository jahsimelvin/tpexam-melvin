package main

import (
	"fmt"
	"sort"
)

func Ft_coin(pieces []int, montant int) int {
	if montant < 0 {
		return -1
	}
	if montant == 0 {
		return 0
	}

	sort.Sort(sort.Reverse(sort.IntSlice(pieces)))

	compteur := 0

	for _, piece := range pieces {
		for montant >= piece {
			montant -= piece
			compteur++
		}
		if montant == 0 {
			break
		}
	}

	if montant > 0 {
		return -1
	}

	return compteur
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // Résultat : 3 (car 11 == 5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // Résultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // Résultat : 0
}
