/*
 * Random number generator based on LCG algorithm
 * reference : https://en.wikipedia.org/wiki/Linear_congruential_generator
 */
package main

import (
	"fmt"
	"flag"
	"time"
)

/*
 * Parameters used in BSD's libc rand()
 */
const (
	BSD_M = 1<<31
	BSD_A = 1103515245
	BSD_C = 12345
)

/*
 * LCG Parameters
 */
type lcgParams struct {
	M int  /* modulus    */
	A int  /* multiplier */
	C int  /* increment  */
}
var parameters lcgParams;

/*
 * Get time based seed
 */
func getseed() int {
	nano := int(time.Now().UnixNano())
	return nano % parameters.M
}

/*
 * Get next value in LCG sequence
 */
func nextVal(x int) int {
	return (parameters.A * x + parameters.C) % parameters.M
}

func main() {
	var x0 int
	var yn string

	/* Process input paremeters */
	flag.IntVar(&parameters.M, "m", BSD_M, "modulus")
	flag.IntVar(&parameters.A, "a", BSD_A, "multipler")
	flag.IntVar(&parameters.C, "c", BSD_C, "increment")
	flag.IntVar(&x0, "x0",  getseed(), "seed")
	flag.Parse()

	x := x0
	fmt.Println("x0 :", x0)	
	for true {
		/* Generate next random number */
		x = nextVal(x)
		fmt.Println("rand() :", x)

		fmt.Print("Continue?[y/n] ")
		fmt.Scanln(&yn)
		if (yn == "n") {
			break
		}
	}
}