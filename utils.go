package main

import (
    "math/rand"
    "time"
)

var Rand *privateRand = &privateRand{rand.New(rand.NewSource(time.Now().UnixNano()))}

type privateRand struct {
    *rand.Rand
}

func (p *privateRand) Number(numberRange ...int) int {
    nr := 0
    if len(numberRange) > 1 {
        nr = 1
        nr = p.Intn(numberRange[1]-numberRange[0]) + numberRange[0]
    } else {
        nr = p.Intn(numberRange[0])
    }
    return nr
}
