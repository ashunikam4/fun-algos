/*
 * Random number generator based on BBS algorithm
 * reference : https://en.wikipedia.org/wiki/Blum_Blum_Shub
 */
package main

import (
    "fmt"
    "flag"
    "os"
)

/*
 * Default Parameters used
 */
const (
    P0 = 3263052707
    Q0 = 5847777359
    X0 = 1033830034
)

/*
 * BBS Parameters
 */
type BBS struct {
    p int /* 1st prime factor */
    q int /* 2nd prime factor */

    state int /* BBS sequence */
}

/*
 * Basic check for primality
 */
 func isPrime(x int) bool {
    if (x == 1) {
        return false
    }

    for i := 2; i * i <= x; i++ {
        if ((x % i) == 0) {
            return false
        }
    }
    return true
}


/*
 * Check for safe primality
 * reference: https://en.wikipedia.org/wiki/Safe_and_Sophie_Germain_primes
 */
func isSafePrime(x int) bool {
    return isPrime(x) && isPrime((x - 1) / 2)
}

/*
 * Get GCD of two numbers (by Euclid's algorithm)
 */
func gcd(a int, b int) int {
    if (a < b) {
        return gcd(b, a)
    } else if (b == 0) {
        return a
    } else {
        return gcd(b, a % b)
    }
}

/*
 * Check if two numbers are coprime
 */
func areCoPrime(a int, b int) bool {
    return (gcd(a, b) == 1)
}

/*
 * Verify Parameters
 */
func (bbs * BBS) verify() bool {
    /*
     * Rule 1: p & q must be safe primes
     */
    if (!isSafePrime(bbs.p) || !isSafePrime(bbs.q)) {
        fmt.Println("Error: Rule 1 violated (p & q must be safe primes)")
        return false
    }

    /*
     * Rule 2: p and q must be congruent to 3 (mod 4)
     */
     if ((bbs.p % 4) != 3 || (bbs.q % 4) != 3) {
        fmt.Println("Error: Rule 2 violated (p and q must be congruent to 3 (mod 4))")
        return false
    }

    /*
     * Rule 3: seed is coprime to p and q
     */
     if (!areCoPrime(bbs.p * bbs.q, bbs.state)) {
        fmt.Println("Error: Rule 3 violated (seed must be coprime with p and q)")
        return false
    }

    return true
}

/*
 * Get next bit
 */
func (bbs *BBS) nextBit() int {
    M := bbs.p * bbs.q
    bbs.state = (bbs.state * bbs.state) % M
    return (bbs.state & 1) /* lsb */
}

/*
 * Get next value
 */
func (bbs *BBS) nextVal() uint32 {
    output := uint32(0)

    /* get 32 bits */
    for i := 0; i < 32; i++ {
        output = (output << 1) | uint32(bbs.nextBit())
    }
    return output
}

/*
 * Ask user if we should terminate
 */
 func askForEnd() bool {
    var yn string
    fmt.Println("Continue?[y/n]")
    fmt.Scanln(&yn)
    return (yn == "n")
}

func main() {
    var bbs BBS

    /* Process input parameters */
    flag.IntVar(&bbs.p, "p", P0, "1st prime factor")
    flag.IntVar(&bbs.q, "q", Q0, "2nd prime factor")
    flag.IntVar(&bbs.state, "x0", X0, "seed")
    flag.Parse()

    /* Verify inputs */
    if (!bbs.verify()) {
        os.Exit(1);
    }

    fmt.Println("x0 :", bbs.state)
    for true {
        /* Generate next random number */
        x := bbs.nextVal()
        fmt.Printf("rand() : %d (0b%b)\n",x, x)

        if (askForEnd()) {
            break
        }
    }
}