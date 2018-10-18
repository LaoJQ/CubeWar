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

    role := cube.roles[dice.selfFace]
    role.Move(cube, point)
    return nil
}

func NewDice() *Dice {
    return &Dice{}
}
