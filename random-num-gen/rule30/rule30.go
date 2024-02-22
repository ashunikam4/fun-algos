/*
 * Random number generator based on Rule30 algorithm
 * reference : https://en.wikipedia.org/wiki/Rule_30#Rule_set
 */
 package main

import (
    "fmt"
    "flag"
    "time"
)

/*
 * Get random seed
 */
func tseed() uint64 {
    t := time.Now()
    return uint64(t.UnixNano())
}

/*
 * Get next state
 * For LSB and MSB, we wrap around i.e. right(LSB) = MSB, left(MSB) = LSB
 */
func evolve(state uint64) uint64 {
    state = (state >> 1 | state << 63) ^ (state | (state << 1 | state >> 63))
    return state
}

/*
 * Print current state using 'o' and '.'
 * 'o' -> 1, '.' -> 1
 */
func printState(state uint64) {
    for i := 0; i < 64; i++ {
        if (((state >> i) & 1) == 0) {
            fmt.Printf(".")
        } else {
            fmt.Printf("o")
        }
    }
    fmt.Println()
}

/*
 * Ask user if we should terminate
 * "n" implies terminate.
 * other inputs implies continue.
 */
func askForEnd() bool {
    var yn string
    fmt.Scanln(&yn)
    return (yn == "n")
}

func main() {
    var state uint64

    /* Process input parameters */
    flag.Uint64Var(&state, "s0", tseed(), "initial state")
    flag.Parse()

    fmt.Println("Press 'n' to end, ENTER to continue")
    for true {
        printState(state)

        if (askForEnd()) {
            break
        }
        state = evolve(state)
    }
}