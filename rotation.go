package main

import (
    "fmt"
    "errors"
)

type Rotation struct {
    // 属性
    *Role
    num int // 剩余道具数

    // 使用参数
    targetFace int // 旋转的面索引
    clockWise bool // 顺时针true, 逆时针false
}

func (rotation *Rotation) Use(cube *Cube) (_ string, err error) {
    if rotation.num <= 0 {
        fmt.Println("Rotation.num =< 0")
        err = errors.New("[ERR] Rotation.num =< 0")
        return
    }
    rotate(cube, rotation)
    rotation.num--
    return
}

func NewRotation() *Rotation {
    return &Rotation{}
}

type RotateRule struct {
    faceIdx int
    gridIdx []int
}

// 顺时针后: idx=0重新刷新, idx=2清空
// 逆时针后: idx=1重新刷新, idx=3清空
var RotateRules [][]RotateRule = [][]RotateRule{
    []RotateRule{
        RotateRule{2, []int{4,3,2}},
        RotateRule{4, []int{4,3,2}},
        RotateRule{3, []int{0,7,6}},
        RotateRule{5, []int{4,3,2}},
    },
    []RotateRule{
        RotateRule{4, []int{0,7,6}},
        RotateRule{2, []int{0,7,6}},
        RotateRule{5, []int{0,7,6}},
        RotateRule{3, []int{4,3,2}},
    },
    []RotateRule{
        RotateRule{4, []int{6,5,4}},
        RotateRule{0, []int{0,7,6}},
        RotateRule{5, []int{2,1,0}},
        RotateRule{1, []int{4,3,2}},
    },
    []RotateRule{
        RotateRule{0, []int{4,3,2}},
        RotateRule{4, []int{2,1,0}},
        RotateRule{1, []int{0,7,6}},
        RotateRule{5, []int{6,5,4}},
    },
    []RotateRule{
        RotateRule{0, []int{2,1,0}},
        RotateRule{2, []int{2,1,0}},
        RotateRule{1, []int{2,1,0}},
        RotateRule{3, []int{2,1,0}},
    },
    []RotateRule{
        RotateRule{2, []int{6,5,4}},
        RotateRule{0, []int{6,5,4}},
        RotateRule{3, []int{6,5,4}},
        RotateRule{1, []int{6,5,4}},
    },
}

func rotate(cube *Cube, rotation *Rotation) {
    rules := RotateRules[rotation.targetFace]
    sideMove, topMove, batteryMove, refreshMove := 3, 6, 1, 0
    if !rotation.clockWise {
        sideMove, topMove, batteryMove, refreshMove = 1, 2, 3, 1
    }
    for i:=0; i<3; i++ {
        cube.face[rules[0].faceIdx][rules[0].gridIdx[i]],
        cube.face[rules[1].faceIdx][rules[1].gridIdx[i]],
        cube.face[rules[2].faceIdx][rules[2].gridIdx[i]],
        cube.face[rules[3].faceIdx][rules[3].gridIdx[i]] =
            cube.face[rules[(0+sideMove)%4].faceIdx][rules[(0+sideMove)%4].gridIdx[i]],
        cube.face[rules[(1+sideMove)%4].faceIdx][rules[(1+sideMove)%4].gridIdx[i]],
        cube.face[rules[(2+sideMove)%4].faceIdx][rules[(2+sideMove)%4].gridIdx[i]],
        cube.face[rules[(3+sideMove)%4].faceIdx][rules[(3+sideMove)%4].gridIdx[i]]
    }
    
    cFace := cube.face[rotation.targetFace]
    cFace[0], cFace[1], cFace[2], cFace[3], cFace[4], cFace[5], cFace[6], cFace[7] =
        cFace[(0+topMove)%8],cFace[(1+topMove)%8],cFace[(2+topMove)%8],cFace[(3+topMove)%8],cFace[(4+topMove)%8],cFace[(5+topMove)%8],cFace[(6+topMove)%8],cFace[(7+topMove)%8]

    if RoleFace(rotation.targetFace) {
        cube.roles[rotation.targetFace].batteryOrien = (cube.roles[rotation.targetFace].batteryOrien + batteryMove)%4
    }

    refreshFace := rules[refreshMove].faceIdx
    for _, refreshGrid := range rules[refreshMove].gridIdx {
        cube.face[refreshFace][refreshGrid].propId = GenProp()
    }
}
