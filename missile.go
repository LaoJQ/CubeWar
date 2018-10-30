package main

import (
    "fmt"
    "errors"
)

const BETTERY_INDEX = 8

type Missile struct {
    // 属性
    *Role
    num int // 剩余道具数

    // 使用参数
    atkGrid int // 目标格子, 0~8, 8打塔
}

func (missile *Missile) Use(cube *Cube) (_ string, err error) {
    selfFace := missile.Role.faceIdx
    targetFace := RotateRules[selfFace][missile.Role.batteryOrien].faceIdx
    if missile.atkGrid == BETTERY_INDEX && !RoleFace(targetFace) {
        return "", errors.New("[ERR] target face is not a role face")
    }
    
    if missile.num <= 0 {
        fmt.Println("Missile.num =< 0")
        err = errors.New("[ERR] Missile.num =< 0")
        return 
    }

    if missile.atkGrid == BETTERY_INDEX {
        cube.roles[targetFace].batteryHP--
    } else {
        borders := BorderMap[targetFace*10+missile.atkGrid]
        for _, idx := range borders {
            cube.face[idx/10][idx%10].propId = PROP_RUINED
        }
        cube.face[targetFace][missile.atkGrid].propId = PROP_RUINED
    }

    missile.num--
    return
}


var BorderMap map[int][]int = map[int][]int{
    0 : []int{22, 44},
    1 : []int{43},
    2 : []int{30, 42},
    3 : []int{37},
    4 : []int{36, 54},
    5 : []int{53},
    6 : []int{24, 52},
    7 : []int{23},

    10 : []int{32, 40},
    11 : []int{47},
    12 : []int{20, 46},
    13 : []int{27},
    14 : []int{26, 50},
    15 : []int{57},
    16 : []int{34, 56},
    17 : []int{33},

    20 : []int{12, 46},
    21 : []int{45},
    22 : []int{0, 44},
    23 : []int{7},
    24 : []int{6, 52},
    25 : []int{51},
    26 : []int{14, 50},
    27 : []int{33},

    30 : []int{2, 42},
    31 : []int{41},
    32 : []int{10, 40},
    33 : []int{17},
    34 : []int{16, 56},
    35 : []int{55},
    36 : []int{4, 54},
    37 : []int{3},

    40 : []int{10, 32},
    41 : []int{31},
    42 : []int{2, 30},
    43 : []int{1},
    44 : []int{0, 22},
    45 : []int{21},
    46 : []int{12, 20},
    47 : []int{11},

    50 : []int{14, 26},
    51 : []int{25},
    52 : []int{6, 44},
    53 : []int{5},
    54 : []int{4, 36},
    55 : []int{35},
    56 : []int{16, 34},
    57 : []int{15},
}
