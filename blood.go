package main

import (
    "fmt"
    "errors"
)

type Blood struct {
    // 属性
    *Role
    num int // 剩余道具数

    // 使用参数
}

func (blood *Blood) Use(cube *Cube) (_ string, err error) {
    if blood.num <= 0 {
        fmt.Println("Blood.num =< 0")
        err = errors.New("[ERR] Blood.num =< 0")
        return
    }
    blood.num--
    blood.batteryHP++
    return
}

