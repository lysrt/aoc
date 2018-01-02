package main

import (
	"fmt"
	"time"
)

// Result := 911
func asm() int {
	start := time.Now()
	input := 109300
	upper := input + 17000

	count := 0
	for current := input; current <= upper; current += 17 {
		fmt.Println(current, "/", upper)
		// label:
		// for factor1 := 2; factor1 != current; factor1++ {
		// 	for factor2 := factor1; factor2 != current; factor2++ {
		// 		if (factor1 * factor2) == current {
		// 			count++
		// 			break label
		// 		}
		// 	}
		// }

		for f := 2; f < current; f++ {
			if current%f == 0 {
				count++
				break
			}
		}
	}
	fmt.Println("Elapsed:", time.Since(start))
	return count
}
