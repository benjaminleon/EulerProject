package main

/*

In the United Kingdom the currency is made up of pound (£) and pence (p).
There are eight coins in general circulation:
1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).

It is possible to make £2 in the following way:
1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/
import (
	"fmt"
)

func main() {

	ways := 0
	for two_pound := 0; two_pound <= 1; two_pound++ {
		for one_pound := 0; one_pound <= 2; one_pound++ {
			for fifty_p := 0; fifty_p <= 4; fifty_p++ {
				for twenty_p := 0; twenty_p <= 10; twenty_p++ {
					for ten_p := 0; ten_p <= 20; ten_p++ {
						for five_p := 0; five_p <= 40; five_p++ {
							for two_p := 0; two_p <= 100; two_p++ {
								for one_p := 0; one_p <= 200; one_p++ {
									sum := 200*two_pound + 100*one_pound + 50*fifty_p + 20*twenty_p + 10*ten_p + 5*five_p + 2*two_p + 1*one_p
									if sum >= 200 {
										if sum == 200 {
											fmt.Printf("%d %d %d %d %d %d %d %d\n",
												two_pound, one_pound, fifty_p, twenty_p, ten_p, five_p, two_p, one_p)
											ways++
										}
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("There are %d ways\n", ways)
}
