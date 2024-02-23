# `Blum Blum Shub`
Invented in 1968, It is provably strong RNG. The algorithm is short and simple:

$x_{i+1}$ = $x_i^2\ mod\ M$

The strength comes from the *quadratic residuosity* problem i.e. Given a composite number $n$, find whether $x$ is a perfect square modulo $n$. It has been proven that this is a NP hard problem.

There are some drawbacks: Firstly, it is computationally expensive. Secondly, there are lot of constraints that seed and M needs to satisfy. 

1. $M\ =\ p.q$ where $p$ and $q$ are safe primes
2. $p$ and $q$ are congruent to 3 (mod 4)
3. $seed$ is coprime to $M$

## Languages used
- `golang`

## References
- `https://en.wikipedia.org/wiki/Blum_Blum_Shub`
- `https://homes.luddy.indiana.edu/kapadia/project2/node11.html`
- `https://www.gkbrk.com/blum-blum-shub`
