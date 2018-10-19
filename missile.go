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

func (missile *Missile) Use(cube *Cube) (_ string, err error) {
    if missile.num <= 0 {
        fmt.Println("Missile.num =< 0")
        err = errors.New("[ERR] Missile.num =< 0")
        return 
    }
    missile.num--

    targetFace := RotateRules[missile.selfFace][cube.roles[missile.selfFace].batteryOrien].faceIdx
    cube.roles[targetFace].batteryHP--
    return
}

func NewMissile() *Missile {
    return &Missile{}
}
