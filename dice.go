package main

import (
    "fmt"
    // "error"
)

type Dice struct {
    // 使用参数
    selfFace int // 面索引, [0,1,2,3,4,5]
}

func (dice *Dice) Use(cube *Cube) error {
    point := Rand.Number(1, 7)
    fmt.Println("dice point: ", point)
    oldSquareIdx := cube.roles[dice.selfFace].squareIdx
    newSquareIdx := (oldSquareIdx+point)%8
    cube.roles[dice.selfFace].squareIdx = newSquareIdx
    return nil
}

func NewDice() *Dice {
    return &Dice{}
}
