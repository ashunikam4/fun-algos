/*
 * Random number generator based on LFSR algorithm
 * reference : https://en.wikipedia.org/wiki/Linear-feedback_shift_register
 */
package main

import (
	"fmt"
	"flag"
	"time"
)

/*
 * Default Parameters used
 */
const (
	N0 = 4
	T0 = 0b10
	K0 = 3
)

/*
 * LFSR Parameters
 */
type LFSR struct {
	N int    /* LFSR size   */
	T int    /* taps  	  	*/
	K int  	 /* output size */

	reg int  /* LFSR values */
}

/*
 * Get time-based random seed
 */
func (lfsr *LFSR) tseed() int {
	t := time.Now()
	nano := int(t.UnixNano())
	return nano % (1 << lfsr.N - 1)
}

/*
 * Get next Bit in LFSR sequence
 */
func (lfsr *LFSR) nextBit() int {
	output := lfsr.reg & 1  /* output rightmost bit */
	lsb := output
	taps := lfsr.T

	for taps > 0 {
		/* Check the register at tap's rightmost set position */
		if ((lfsr.reg & (taps & -taps)) > 0) {
			lsb ^= 1  /* LSB is XOR of tap positions and outputs */
		}

		/* unset the tap's rightmost set position */
		taps = taps & (taps - 1)
	}

	/* Shift the register to right by 1, add lsb to left */
	lfsr.reg = (lfsr.reg >> 1)
	lfsr.reg |= lsb << (lfsr.N - 1)
	return output
}

/*
 * Get next value in LFSR sequence
 */
func (lfsr *LFSR) nextVal() int {
	output := 0

	/* Pop next K bits */
	for i := 0; i < lfsr.K; i++ {
		output = (output << 1) | lfsr.nextBit()
	}
	return output
}

func main() {
	var lfsr LFSR
	var yn string

	/* Process input parameters */
	flag.IntVar(&lfsr.N, "n", N0, "LFSR width")
	flag.IntVar(&lfsr.T, "t", T0, "taps")
	flag.IntVar(&lfsr.K, "k", K0, "random number width")
	flag.IntVar(&lfsr.reg, "x0", lfsr.tseed(), "seed")
	flag.Parse()

	fmt.Printf("x0 : 0b%b\n", lfsr.reg)
	for true {
		/* Generate next random number */
		x := lfsr.nextVal()
		fmt.Printf("rand(), lfsr.reg : 0b%b, 0b%b\n", x, lfsr.reg)

		fmt.Print("Continue?[y/n] ")
		fmt.Scanln(&yn)
		if (yn == "n") {
			break
		}
	}
}