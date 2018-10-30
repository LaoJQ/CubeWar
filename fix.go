package main

import (
    "fmt"
    "errors"
)

type Fix struct {
    // 属性
    *Role
    num int // 剩余道具数

    // 使用参数
    target int // face*10+grid
}

func (fix *Fix) Use(cube *Cube) (_ string, err error) {
    if fix.num <= 0 {
        fmt.Println("Fix.num =< 0")
        err = errors.New("[ERR] Fix.num =< 0")
        return
    }
    if fix.target%10 == BETTERY_INDEX {
        fmt.Println("Grid can not be bettery")
        err = errors.New("Grid can not be bettery")
        return
    }
    if cube.face[fix.target/10][fix.target%10].propId != PROP_RUINED {
        fmt.Println("This grid is not ruined")
        err = errors.New("This grid is not ruined")
        return
    }

    borders := BorderMap[fix.target]
    for _, idx := range borders {
        cube.face[idx/10][idx%10].propId = GenProp()
    }

    fix.num--
    return
}


