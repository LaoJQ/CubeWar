package main

import (
    "fmt"
    "errors"
)

type Missile struct {
    // 属性
    *Role
    num int // 剩余道具数

    // 使用参数
}

func (missile *Missile) Use(cube *Cube) (_ string, err error) {
    selfFace := missile.Role.faceIdx
    targetFace := RotateRules[selfFace][missile.Role.batteryOrien].faceIdx
    if !RoleFace(targetFace) {
        return "", errors.New("[ERR] target face is not a role face")
    }
    
    if missile.num <= 0 {
        fmt.Println("Missile.num =< 0")
        err = errors.New("[ERR] Missile.num =< 0")
        return 
    }
    missile.num--
    cube.roles[targetFace].batteryHP--
    return
}

