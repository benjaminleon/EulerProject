package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := map[string]bool{}
	var result big.Int // Use this to access the power (exp) method

	for a := int64(2); a <= 100; a++ {
		for b := int64(2); b <= 100; b++ {
			m[result.Exp(big.NewInt(a), big.NewInt(b), nil).String()] = true
		}
	}

	fmt.Printf("There are %d distinct powers\n", len(m))
}
