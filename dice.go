package main

import (
    "fmt"
    "errors"
)

type Dice struct {
    // 属性
    *Role
    
    // 使用参数
}

func (dice *Dice) Use(cube *Cube) (string, error) {
    if !RoleFace(dice.Role.faceIdx) {
        return "", errors.New("[ERR] handle face is not a role face")
    }
    point := Rand.Number(1, 7)
    dice.Role.Move(cube, point)
    return fmt.Sprintln("dice point: ", point), nil
}

func NewDice() *Dice {
    return &Dice{}
}
