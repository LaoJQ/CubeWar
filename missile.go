package main

import (
    "fmt"
    "errors"
)

type Missile struct {
    // 属性
    num int // 剩余道具数

    // 使用参数
    selfFace int // 所在的面, [0,1,2,3,4,5]
}

func (missile *Missile) Use(cube *Cube) error {
    if missile.num <= 0 {
        fmt.Println("Rotation.num =< 0")
        return errors.New("Rotation.num =< 0")
    }
    missile.num--

    targetFace := RotateRules[missile.selfFace][cube.face[missile.selfFace][8].batteryOrien].faceIdx
    cube.face[targetFace][8].batteryHP--
    return nil
}

func NewMissile() *Missile {
    return &Missile{}
}
