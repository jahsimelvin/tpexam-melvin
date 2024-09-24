package main

import (
	"fmt"
)


func Ft_profit(prix []int) int {
	if len(prix) == 0 {
		return 0
	}

	prixMin := prix[0]
	beneficeMax := 0

	for _, valeur := range prix {
		if valeur < prixMin {
			prixMin = valeur
		} else {
			bénéfice := valeur - prixMin
			if bénéfice > beneficeMax {
			 beneficeMax = bénéfice
			}
		}
	}

	return beneficeMax
}

func main() {
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // Résultat : 5
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))    // Résultat : 0
}
