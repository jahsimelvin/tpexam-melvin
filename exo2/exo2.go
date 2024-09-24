package main

import (
	"fmt"
)

func Ft_missing(nombres []int) int {
	n := len(nombres)
	present := make([]bool, n+1)

	for _, nombre := range nombres {
		if nombre <= n {
			present[nombre] = true
		}
	}

	for i := 0; i <= n; i++ {
		if !present[i] {
			return i
		}
	}

	return n
}

func main() {
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // Résultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // Résultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Résultat : 8
}
